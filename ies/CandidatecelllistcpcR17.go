package ies

// CandidateCellListCPC-r17 ::= SEQUENCE OF CandidateCellCPC-r17
// SIZE (1..maxFreq)
type CandidatecelllistcpcR17 struct {
	Value []CandidatecellcpcR17 `lb:1,ub:maxFreq`
}
