package ies

// SchedulingInfoList ::= SEQUENCE OF SchedulingInfo
// SIZE (1..maxSI-Message)
type Schedulinginfolist struct {
	Value []Schedulinginfo `lb:1,ub:maxSIMessage`
}
