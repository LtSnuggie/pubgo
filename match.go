package pubgo

func (mr *MatchResponse) GetStatsByName() (s map[string]*MatchStats) {
	s = make(map[string]*MatchStats)
	for _, p := range mr.Participants {
		s[p.Attributes.Stats.Name] = &p.Attributes.Stats
	}
	return
}

func (mr *MatchResponse) GetStatsByWinRank() (s [][]*MatchStats) {
	s = make([][]*MatchStats, len(mr.Participants))
	for _, p := range mr.Participants {
		stat := p.Attributes.Stats
		v := s[stat.WinPlace-1]
		v = append(v, &stat)
		s[stat.WinPlace-1] = v
		// s[p.Attributes.Stats.Name] = &p.Attributes.Stats
	}
	return
}
