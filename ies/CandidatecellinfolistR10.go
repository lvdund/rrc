package ies

// CandidateCellInfoList-r10 ::= SEQUENCE OF CandidateCellInfo-r10
// SIZE (1..maxFreq)
type CandidatecellinfolistR10 struct {
	Value []CandidatecellinfoR10 `lb:1,ub:maxFreq`
}
