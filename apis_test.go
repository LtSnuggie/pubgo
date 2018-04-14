package pubgo_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/ltsnuggie/pubgo"
)

type Conf struct {
	Key       string `json:"key"`
	RateLimit int    `json:"rateLimit"`
}

const (
	expectedAPI = "v8.4.0"
)

var session *pubgo.Session
var players = []string{"Lt Snuggie", "DinnerPlates7", "WARISTHEANSWER0"}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestGetStatus(t *testing.T) {
	size := session.GetStatus(func(sr pubgo.StatusResponse, err error) {
		if sr.Data.Attributes.Version != expectedAPI {
			t.Errorf("expected version %s but received %s", expectedAPI, sr.Data.Attributes.Version)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetPlayer(t *testing.T) {
	size := session.GetPlayer(players[0], func(prd pubgo.PlayerResponseData, err error) {
		if prd.Attributes.Name != players[0] {
			t.Errorf("expected a player name of %s but received %s", players[0], prd.Attributes.Name)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetPlayers(t *testing.T) {
	size := session.GetPlayers(players, func(pr pubgo.PlayerResponse, err error) {
		if len(players) != len(pr.Data) {
			t.Errorf("expected %d players in PlayerResponse but received %d", len(players), len(pr.Data))
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetMatch(t *testing.T) {
	var id, player string
	session.GetPlayers(players, func(pr pubgo.PlayerResponse, err error) {
		for _, prd := range pr.Data {
			if len(prd.GetMatchIDs()) > 0 {
				id = prd.GetMatchIDs()[0]
				player = prd.Attributes.Name
				return
			}
		}
	})
	if id == "" {
		t.Errorf("no match ids for players %s", strings.Join(players, ","))
		return
	}
	size := session.GetMatch(id, func(mr pubgo.MatchResponse, err error) {
		if mr.GetStatsByName()[player] == nil {
			t.Errorf("expected player %s to be in match but was not found", player)
		}
	})
	if size != 0 {
		t.Errorf("expected a queue size of 0 but received %d", size)
	}
}
func TestGetTelemetry(t *testing.T)          {}
func TestReadTelemetryFromFile(t *testing.T) {}

//Do last to prevent delaying prior tests
func TestGetQueueSize(t *testing.T) {}

func setup() {
	b, _ := ioutil.ReadFile("testdata/conf.js")
	var conf Conf
	json.Unmarshal(b, &conf)
	session, _ = pubgo.New(conf.Key, conf.RateLimit)
}
func teardown() {}
