package ies

// SchedulingInfoListExt-r12 ::= SEQUENCE OF SchedulingInfoExt-r12
// SIZE (1..maxSI-Message)
type SchedulinginfolistextR12 struct {
	Value []SchedulinginfoextR12 `lb:1,ub:maxSIMessage`
}
