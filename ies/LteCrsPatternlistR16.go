package ies

// LTE-CRS-PatternList-r16 ::= SEQUENCE OF RateMatchPatternLTE-CRS
// SIZE (1..maxLTE-CRS-Patterns-r16)
type LteCrsPatternlistR16 struct {
	Value []RatematchpatternlteCrs `lb:1,ub:maxLTECrsPatternsR16`
}
