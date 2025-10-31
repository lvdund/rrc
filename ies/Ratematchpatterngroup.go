package ies

// RateMatchPatternGroup ::= SEQUENCE OF RateMatchPatternGroup-Item
// SIZE (1..maxNrofRateMatchPatternsPerGroup)
type Ratematchpatterngroup struct {
	Value []RatematchpatterngroupItem `lb:1,ub:maxNrofRateMatchPatternsPerGroup`
}
