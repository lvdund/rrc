package ies

// UE-RadioPagingInfo-r17 ::= SEQUENCE
// Extensible
type UeRadiopaginginfoR17 struct {
	PeiSubgroupingsupportbandlistR17 *[]Freqbandindicatornr `lb:1,ub:maxBands`
}
