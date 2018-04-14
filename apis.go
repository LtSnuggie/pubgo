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
	base    = "https://api.playbattlegrounds.com" // base URL for making API calls
	shards  = "/shards/"                          // shards path segment
	matches = "/matches/"                         // matches end point
	players = "/players"                          //players end point
	status  = "/status"                           // status end point

	// XboxAsia - Xbox Asia Region
	XboxAsia = "xbox-as"
	// XboxEurope - Xbox Europe Region
	XboxEurope = "xbox-eu"
	// XboxNorthAmerica - Xbox North America Region
	XboxNorthAmerica = "xbox-na"
	// XboxOceania - Xbox Oceana Region
	XboxOceania = "xbox-oc"
	// PCAsia - PC Asia  Region
	PCAsia = "xbox-as"
	// PCEurope - PC Europe Region
	PCEurope = "xbox-eu"
	// PCNorthAmerica - PC North America Region
	PCNorthAmerica = "xbox-na"
	// PCOceania - PC Oceania Region
	PCOceania = "xbox-oc"
	// PCKoreaJapan - PC Korea/Japan Region
	PCKoreaJapan = "pc-krjp"
	// PCKorea - PC Korea Region
	PCKorea = "pc-kr"
	// PCJapan - PC Japan Region
	PCJapan = "pc-jp"
	// PCKAKAO - PC KAKAO Region
	PCKAKAO = "pc-kakao"
	// PCSouthEastAsia - PC South East Asia Region
	PCSouthEastAsia = "pc-sea"
	// PCSouthAsia - PC South Asia Region
	PCSouthAsia = "pc-sa"
)

// GetQueueSize returns the current size of the poller queue.
// This is useful if implementing additional request limiting.
func (s *Session) GetQueueSize() (size int) {
	size = len(s.poller.queue)
	return
}

// GetStatus retrieves status data from the PUBG servers and passes the StatusResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
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

// GetPlayer retrieves data for the specified player and passes the PlayerResponseData into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
func (s *Session) GetPlayer(name string, clbk func(PlayerResponseData, error)) (size int) {
	s.GetPlayers([]string{name}, func(pr PlayerResponse, err error) {
		if len(pr.Data) > 0 {
			clbk(pr.Data[0], err)
		}
	})
	return s.GetQueueSize()
}

// GetPlayers retrieves data for the passed names and passes the PlayerResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
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

// GetMatch retrieves the match data for a specified match id and passes the MatchResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
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

// GetTelemetry retrieves the telemetry data at a specified url and passes the TelemetryResponse into the given callback.
// Upon retrieval of data the callback passed in is executed. Additionally the size of the
// poller buffer is returned.
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

// ReadTelemetryFromFile parses json telemetry data from a given file
// and returns a TelemetryResponse struct. It is more performant to cache
// telemetry data for future use.
func (s *Session) ReadTelemetryFromFile(path string) (tr TelemetryResponse, err error) {
	var b []byte
	b, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	return parseTelemetry(b)
}

// parseTelemetry reads the telemetry event type from the json
// and passes it to the unmarshaller
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
