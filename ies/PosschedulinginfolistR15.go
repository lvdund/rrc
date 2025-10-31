package ies

// PosSchedulingInfoList-r15 ::= SEQUENCE OF PosSchedulingInfo-r15
// SIZE (1..maxSI-Message)
type PosschedulinginfolistR15 struct {
	Value []PosschedulinginfoR15 `lb:1,ub:maxSIMessage`
}
