package pubgo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	base    = "https://api.playbattlegrounds.com"
	shards  = "/shards/"
	matches = "/matches/"
	players = "/players"
	status  = "/status"

	XboxAsia         = "xbox-as"
	XboxEurope       = "xbox-eu"
	XboxNorthAmerica = "xbox-na"
	XboxOceania      = "xbox-oc"
	PCAsia           = "xbox-as"
	PCEurope         = "xbox-eu"
	PCNorthAmerica   = "xbox-na"
	PCOceania        = "xbox-oc"
	PCKoreaJapan     = "pc-krjp"
	PCKAKAO          = "pc-kakao"
	PCSouthEastAsia  = "pc-sea"
	PCSouthAsia      = "pc-sa"
)

func (s *Session) GetStatus() (sr StatusResponse) {
	req, _ := http.NewRequest("GET", base+status, nil)
	res, _ := s.Client.Do(req)
	var buffer bytes.Buffer
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &sr)
	return
}

func (s *Session) GetMatch(id string) (mr MatchResponse) {
	req, _ := http.NewRequest("GET", base+shards+matches+id, nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	res, _ := s.Client.Do(req)
	var buffer bytes.Buffer
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &mr)
	for _, inc := range mr.Included {
		var check map[string]string
		json.Unmarshal(inc, &check)
		switch check["type"] {
		case "participant":
			var p MatchParticipant
			json.Unmarshal(inc, &p)
			mr.Participants = append(mr.Participants, p)
		case "asset":
			var a MatchAsset
			json.Unmarshal(inc, &a)
			mr.Assets = append(mr.Assets, a)
		case "roster":
			var r MatchRoster
			json.Unmarshal(inc, &r)
			mr.Rosters = append(mr.Rosters, r)
		}
	}
	return
}

func (s *Session) GetPlayer(name string) (prd PlayerResponseData) {
	pr := s.GetPlayers([]string{name})
	if len(pr.Data) > 0 {
		prd = pr.Data[0]
	}
	return
}

func (s *Session) GetPlayers(names []string) (prd PlayerResponse) {
	query := strings.Replace(strings.Join(names, ","), " ", "%20", -1)
	u, _ := url.ParseRequestURI(base + shards + s.region + players + "?filter[playerNames]=" + query)
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	res, _ := s.Client.Do(req)
	var buffer bytes.Buffer
	var pr PlayerResponse
	buffer.ReadFrom(res.Body)
	json.Unmarshal(buffer.Bytes(), &pr)
	prd = pr
	return
}

func (s *Session) GetTelemetry(url string) (tr TelemetryResponse) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	res, _ := s.Client.Do(req)
	var buffer bytes.Buffer
	// var pretty bytes.Buffer
	buffer.ReadFrom(res.Body)
	// data := fmt.Sprint(string(buffer.Bytes()))
	// err := json.Indent(&pretty, buffer.Bytes(), "", "\t")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// ioutil.WriteFile("tele.js", pretty.Bytes(), 0777)
	var v []json.RawMessage
	json.Unmarshal(buffer.Bytes(), &v)
	for _, bts := range v {
		var eval map[string]interface{}
		json.Unmarshal(bts, &eval)
		tr.unmarshalEvent(bts, eval["_T"].(string))
	}
	return
}
