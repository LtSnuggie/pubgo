package pubgo

import (
	"net/http"
	"time"
)

type Session struct {
	// Client    *http.Client
	apiKey string
	region string
	poller *poller
}

func New(key string, rateLimit int) (s *Session, err error) {
	s = &Session{
		// Client:    &http.Client{Timeout: (20 * time.Second)},
		apiKey: key,
		region: XboxNorthAmerica,
		poller: newPoller(&http.Client{Timeout: (20 * time.Second)}, rateLimit),
	}
	return
}
