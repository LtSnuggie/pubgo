package pubgo

func (p *PlayerResponse) GetPlayerData(name string) (prd PlayerResponseData) {
	for _, d := range p.Data {
		if d.Attributes.Name == name {
			prd = d
			return
		}
	}
	return
}

func (prd *PlayerResponseData) GetMatches() (m []MatchData) {
	m = prd.Relationships.Matches.Data
	return
}

func (prd *PlayerResponseData) GetMatchIDs() (ids []string) {
	for _, d := range prd.Relationships.Matches.Data {
		ids = append(ids, d.ID)
	}
	return ids
}
