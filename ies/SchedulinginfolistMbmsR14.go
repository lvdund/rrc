package ies

// SchedulingInfoList-MBMS-r14 ::= SEQUENCE OF SchedulingInfo-MBMS-r14
// SIZE (1..maxSI-Message)
type SchedulinginfolistMbmsR14 struct {
	Value []SchedulinginfoMbmsR14 `lb:1,ub:maxSIMessage`
}
