package ies

// BandCombinationInfoList ::= SEQUENCE OF BandCombinationInfo
// SIZE (1..maxBandComb)
type Bandcombinationinfolist struct {
	Value []Bandcombinationinfo `lb:1,ub:maxBandComb`
}
