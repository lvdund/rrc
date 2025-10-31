package ies

// CandidateCellInfoListCPC-r17 ::= SEQUENCE OF CandidateCellInfo-r17
// SIZE (1..maxFreq)
type CandidatecellinfolistcpcR17 struct {
	Value []CandidatecellinfoR17 `lb:1,ub:maxFreq`
}
