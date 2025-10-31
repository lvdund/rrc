package ies

// GWUS-ProbThreshList-NB-r16 ::= SEQUENCE OF GWUS-Paging-ProbThresh-NB-r16
// SIZE (1..maxGWUS-ProbThresholds-NB-r16)
type GwusProbthreshlistNbR16 struct {
	Value []GwusPagingProbthreshNbR16 `lb:1,ub:maxGWUSProbthresholdsNbR16`
}
