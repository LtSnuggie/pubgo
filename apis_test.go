package pubgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Conf struct {
	Key       string `json:"key"`
	RateLimit int    `json:"rateLimit"`
}

const (
	expectedAPI = "v9.11.2"
)

var session *Session
var testPlayers = []string{"badshroud"}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestGetStatus(t *testing.T) {
	size := session.GetStatus(func(sr StatusResponse, err error) {
		// if sr.Data.Attributes.Version != expectedAPI {
		// 	t.Errorf("expected version %s but received %s", expectedAPI, sr.Data.Attributes.Version)
		// }
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetPlayer(t *testing.T) {
	size := session.GetPlayer(testPlayers[0], PCNorthAmerica, func(prd PlayerResponseData, err error) {
		if prd.Attributes.Name != testPlayers[0] {
			t.Errorf("expected a player name of %s but received %s", testPlayers[0], prd.Attributes.Name)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetPlayersByName(t *testing.T) {
	size := session.GetPlayersByName(testPlayers, PCNorthAmerica, func(pr PlayerResponse, err error) {
		if len(testPlayers) != len(pr.Data) {
			t.Errorf("expected %d players in PlayerResponse but received %d", len(players), len(pr.Data))
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetMatch(t *testing.T) {
	var id, player string
	size := session.GetPlayersByName(testPlayers, PCNorthAmerica, func(pr PlayerResponse, err error) {
		for _, prd := range pr.Data {
			if len(prd.GetMatchIDs()) > 0 {
				id = prd.GetMatchIDs()[0]
				player = prd.Attributes.Name
				return
			}
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
	if id == "" {
		t.Errorf("no match ids for players %s", strings.Join(testPlayers, ","))
		return
	}
	size = session.GetMatch(id, PCNorthAmerica, func(mr MatchResponse, err error) {
		if mr.GetStatsByName()[player] == nil {
			t.Errorf("expected player %s to be in match but was not found", player)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}

func TestGetSamples(t *testing.T) {
	size := session.GetSampleMatches(PCNorthAmerica, func(sr SamplesResponse, err error) {
		if err != nil {
			t.Errorf("API returned error: %s", err.Error())
		}
		if len(sr.GetMatches()) == 0 {
			t.Errorf("Expected samples in SampleResponse but received 0")
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}

func TestGetSeasons(t *testing.T)            {}
func TestGetTelemetry(t *testing.T)          {}
func TestReadTelemetryFromFile(t *testing.T) {}

//Do last to prevent delaying prior tests
func TestGetQueueSize(t *testing.T) {}

func setup() {
	b, _ := ioutil.ReadFile("conf.json")
	var conf Conf
	json.Unmarshal(b, &conf)
	var err error
	session, err = New(conf.Key, conf.RateLimit)
	if err != nil {
		panic("Error creating API Client")
	}
}
func teardown() {}
