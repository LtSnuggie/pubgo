package pubgo

// GetStatsByName is a helper function to retrieve player MatchStats
// from a MatchResponse. A map is more performant than a slice in
// larger data sets so it is recommended to use this map instead of
// iterating through the MatchResponse data looking for players
func (mr *MatchResponse) GetStatsByName() (s map[string]*MatchStats) {
	s = make(map[string]*MatchStats)
	for _, p := range mr.Participants {
		s[p.Attributes.Stats.Name] = &p.Attributes.Stats
	}
	return
}

// GetStatsByWinRank is a helper function to retrieve player MatchStats
//  by final rank from a MatchResponse. A map is more performant than a
// slice in larger data sets so it is recommended to use this map instead
// of iterating through the MatchResponse data looking for player ranks
func (mr *MatchResponse) GetStatsByWinRank() (s map[int][]*MatchStats) {
	s = make(map[int][]*MatchStats)
	for _, p := range mr.Participants {
		stat := p.Attributes.Stats
		v := s[stat.WinPlace-1]
		v = append(v, &stat)
		s[stat.WinPlace-1] = v
		// s[p.Attributes.Stats.Name] = &p.Attributes.Stats
	}
	return
}

func (mr *MatchResponse) GetMatchID() (id string) {
	return mr.Data.ID
}
