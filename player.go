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

func (p *PlayerResponseData) GetMatches() (m []MatchData) {
	m = p.Relationships.Matches.Data
	return
}
