// Package pubgo is a wrapper with helper functions for accessing pubg
// servers. Only thing that is required is a developer API key.
package pubgo

import (
	"net/http"
	"time"
)

// Session is the main struct for pubgo
type Session struct {
	apiKey string  // developers api key, used to make calls
	region string  // player region to use when requesting data from server
	poller *poller // poller is responsible for executing the requests as well as maintaining rate limit queue
}

// New returns a new defaulted Session struct.
func New(key string, rateLimit int) (s *Session, err error) {
	s = &Session{
		apiKey: key,
		region: XboxNorthAmerica,
		poller: newPoller(&http.Client{Timeout: (20 * time.Second)}, rateLimit),
	}
	return
}

func GetShards() (shards []string) {
	shards = append(shards, XboxNorthAmerica)
	shards = append(shards, XboxOceania)
	shards = append(shards, XboxEurope)
	shards = append(shards, XboxAsia)
	shards = append(shards, PCAsia)
	shards = append(shards, PCKAKAO)
	shards = append(shards, PCKorea)
	shards = append(shards, PCJapan)
	shards = append(shards, PCEurope)
	shards = append(shards, PCOceania)
	shards = append(shards, PCSouthAsia)
	shards = append(shards, PCKoreaJapan)
	shards = append(shards, PCNorthAmerica)
	return
}
