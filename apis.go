package pubgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	PCKorea          = "pc-kr"
	PCJapan          = "pc-jp"
	PCKAKAO          = "pc-kakao"
	PCSouthEastAsia  = "pc-sea"
	PCSouthAsia      = "pc-sa"
)

func (s *Session) GetQueueSize() (size int) {
	size = len(s.poller.queue)
	return
}

func (s *Session) GetStatus(clbk func(StatusResponse, error)) (size int) {
	req, _ := http.NewRequest("GET", base+status, nil)
	s.poller.Request(req, func(res *http.Response, err error) {
		var sr StatusResponse
		if err != nil {
			clbk(sr, err)
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &sr)
		clbk(sr, err)
	})
	return s.GetQueueSize()
}

func (s *Session) GetPlayer(name string, clbk func(PlayerResponseData, error)) (size int) {
	s.GetPlayers([]string{name}, func(pr PlayerResponse, err error) {
		if len(pr.Data) > 0 {
			clbk(pr.Data[0], err)
		}
	})
	return s.GetQueueSize()
}

func (s *Session) GetPlayers(names []string, clbk func(PlayerResponse, error)) (size int) {
	query := strings.Replace(strings.Join(names, ","), " ", "%20", -1)
	u, _ := url.ParseRequestURI(base + shards + s.region + players + "?filter[playerNames]=" + query)
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	s.poller.Request(req, func(res *http.Response, err error) {
		var pr PlayerResponse
		if err != nil {
			clbk(pr, err)
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &pr)
		clbk(pr, err)
	})
	return s.GetQueueSize()
}

func (s *Session) GetMatch(id string, clbk func(MatchResponse, error)) (size int) {
	req, _ := http.NewRequest("GET", base+shards+matches+id, nil)
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	s.poller.Request(req, func(res *http.Response, err error) {
		var mr MatchResponse
		if err != nil {
			clbk(mr, err)
		}
		var buffer bytes.Buffer
		buffer.ReadFrom(res.Body)
		err = json.Unmarshal(buffer.Bytes(), &mr)
		if err != nil {
			clbk(mr, err)
		}
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
		clbk(mr, nil)
	})
	return s.GetQueueSize()
}

func (s *Session) GetTelemetry(url string, clbk func(TelemetryResponse, error)) (size int) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	s.poller.Request(req, func(res *http.Response, err error) {
		if err != nil {
			clbk(TelemetryResponse{}, err)
			return
		}
		var buffer bytes.Buffer
		// var pretty bytes.Buffer
		buffer.ReadFrom(res.Body)
		// data := fmt.Sprint(string(buffer.Bytes()))
		// err := json.Indent(&pretty, buffer.Bytes(), "", "\t")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		// ioutil.WriteFile("tele.js", pretty.Bytes(), 0777)
		clbk(parseTelemetry(buffer.Bytes()))
	})
	return s.GetQueueSize()
}

func (s *Session) ReadTelemetryFromFile(path string) (tr TelemetryResponse, err error) {
	var b []byte
	b, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	return parseTelemetry(b)
}

func parseTelemetry(b []byte) (tr TelemetryResponse, err error) {
	var v []json.RawMessage
	json.Unmarshal(b, &v)
	for _, bts := range v {
		var eval map[string]interface{}
		err = json.Unmarshal(bts, &eval)
		if err != nil {
			return
		}
		tr.unmarshalEvent(bts, eval["_T"].(string))
	}
	return
}
