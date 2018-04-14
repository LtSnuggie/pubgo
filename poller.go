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
	count  int
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
		0,
		client,
	}
}

func (p *poller) Request(req *http.Request, clbk func(*http.Response, error)) {
	r := request{
		req,
		clbk,
	}
	p.Lock()
	defer p.Unlock()
	if p.count < p.limit {
		if !p.exec(r) {
			p.queue = append(p.queue, r)
			go func() {
				time.Sleep(60 * time.Second)
				p.decrement()
			}()
		}
	} else {
		p.queue = append(p.queue, r)
	}
}

func (p *poller) exec(r request) (ok bool) {
	p.count++
	res, err := p.client.Do(r.req)
	if res.StatusCode != limitReached {
		r.clbk(res, err)
		go func() {
			time.Sleep(60 * time.Second)
			p.decrement()
		}()
		ok = true
	}
	return
}

func (p *poller) decrement() {
	p.Lock()
	defer p.Unlock()
	p.count--
	if len(p.queue) > 0 {
		p.exec(p.queue[0])
		p.queue = p.queue[1:]
	}
}
