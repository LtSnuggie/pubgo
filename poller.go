package pubgo

import (
	"net/http"
	"sync"
	"time"
)

const (
	limitReached = 429
)

type poller struct {
	*sync.Mutex
	queue  []request
	limit  int
	client *http.Client
}

type request struct {
	req  *http.Request
	clbk func(*http.Response, error)
}

func newPoller(client *http.Client, limit int) *poller {
	return &poller{
		&sync.Mutex{},
		make([]request, 0),
		limit,
		client,
	}
}

func (p *poller) Request(req *http.Request, clbk func(*http.Response, error)) {
	r := request{
		req,
		clbk,
	}
	if len(p.queue) > 0 {
		p.Lock()
		p.queue = append(p.queue, r)
		p.Unlock()
		go func() {
			time.Sleep(60 * time.Second)
			p.poll()
		}()
	} else {
		p.exec(r)
	}
}

func (p *poller) exec(r request) (remove bool) {
	res, err := p.client.Do(r.req)
	remove = true
	if err == nil {
		switch res.StatusCode {
		case 200:
		case 401:
			err = NewInvalidKeyError(r.req.URL.String())
		case 404:
			err = NewNotFoundError(r.req.URL.String())
		case 415:
			err = NewIncorrectContentTypeError(r.req.URL.String())
		case 429:
			err = NewTooManyRequestsError(r.req.URL.String())
			// Hit the limit, put this request back in queue - Should be end?
			p.Lock()
			p.queue = append(p.queue, r)
			p.Unlock()
		default:
			err = NewUnhandledStatusCodeError(r.req.URL.String(), res.Status)
		}
	} else {
		p.Lock()
		p.queue = append(p.queue, r)
		p.Unlock()
	}
	r.clbk(res, err)
	go func() {
		time.Sleep(60 * time.Second)
		p.poll()
	}()
	return
}

func (p *poller) poll() {
	p.Lock()
	defer p.Unlock()
	if len(p.queue) > 0 {
		go p.exec(p.queue[0])
		p.queue = p.queue[1:]
	}
}
