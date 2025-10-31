package ies

// SchedulingInfoList-BR-r13 ::= SEQUENCE OF SchedulingInfo-BR-r13
// SIZE (1..maxSI-Message)
type SchedulinginfolistBrR13 struct {
	Value []SchedulinginfoBrR13 `lb:1,ub:maxSIMessage`
}
