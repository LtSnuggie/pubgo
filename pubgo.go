package pubgo

import (
	"net/http"
	"time"
)

type Session struct {
	Client *http.Client
	apiKey string
	region string
}

func New(key string) (s *Session, err error) {
	s = &Session{
		Client: &http.Client{Timeout: (20 * time.Second)},
		apiKey: key,
		region: XboxNorthAmerica,
	}
	return
}
