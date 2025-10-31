package ies

// SchedulingInfoList-NB-r13 ::= SEQUENCE OF SchedulingInfo-NB-r13
// SIZE (1..maxSI-Message-NB-r13)
type SchedulinginfolistNbR13 struct {
	Value []SchedulinginfoNbR13 `lb:1,ub:maxSIMessageNbR13`
}
