package pubgo

// GetMatches is a helper function to extract a slice of MatchData
// from SamplesResponse. This is useful if the caller is only
// interested in MatchData
func (sr *SamplesResponse) GetMatches() (m []MatchData) {
	m = sr.Data.Relationships.Matches.Data
	return
}
