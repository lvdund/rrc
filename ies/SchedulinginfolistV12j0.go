package ies

// SchedulingInfoList-v12j0 ::= SEQUENCE OF SchedulingInfo-v12j0
// SIZE (1..maxSI-Message)
type SchedulinginfolistV12j0 struct {
	Value []SchedulinginfoV12j0 `lb:1,ub:maxSIMessage`
}
