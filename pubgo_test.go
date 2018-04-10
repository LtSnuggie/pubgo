package pubgo_test

import (
	"encoding/json"
	"io/ioutil"
	"pubgo"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type Conf struct {
	Key string `json:"key"`
}

func TestX(t *testing.T) {
	b, _ := ioutil.ReadFile("testdata/conf.js")
	var conf Conf
	json.Unmarshal(b, &conf)
	s, _ := pubgo.New(conf.Key)
	p := s.GetPlayers([]string{"Lt Snuggie", "DinnerPlates7", "WARISTHEANSWER0"})
	spew.Dump(p.Data[0].Attributes)
	mr := s.GetMatch(p.Data[2].Relationships.Matches.Data[0].ID)
	for _, part := range mr.Participants {
		if part.Attributes.Stats.WinPlace < 10 {
			spew.Dump(part.Attributes.Stats)
		}
	}
	// for _, asset := range mr.Assets {
	// 	tr := s.GetTelemetry(asset.Attributes.URL)
	// 	for _, e := range tr.PlayerTakeDamageEvents {
	// 		if e.Victim.Name == "WARISTHEANSWER0" {
	// 			fmt.Printf("%s took %v damage via %s from %s\n", e.Victim.Name, e.Damage, e.DamageCauserName, e.Attacker.Name)
	// 		}
	// 	}
	// }
}
