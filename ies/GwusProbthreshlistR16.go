package ies

// GWUS-ProbThreshList-r16 ::= SEQUENCE OF GWUS-PagingProbThresh-r16
// SIZE (1..maxGWUS-ProbThresholds-r16)
type GwusProbthreshlistR16 struct {
	Value []GwusPagingprobthreshR16 `lb:1,ub:maxGWUSProbthresholdsR16`
}
