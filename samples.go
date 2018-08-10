package pubgo

// GetMatches is a helper function to extract a slice of MatchData
// from SamplesResponseData. This is useful if the caller is only
// interested in MatchData
func (srd *SamplesResponseData) GetMatches() (m []MatchData) {
	m = srd.Relationships.Matches.Data
	return
}
