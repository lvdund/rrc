package ies

// SupportedBandListNR-r15 ::= SEQUENCE OF SupportedBandNR-r15
// SIZE (1..maxBandsNR-r15)
type SupportedbandlistnrR15 struct {
	Value []SupportedbandnrR15 `lb:1,ub:maxBandsNRR15`
}
