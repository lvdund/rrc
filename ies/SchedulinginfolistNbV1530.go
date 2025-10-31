package ies

// SchedulingInfoList-NB-v1530 ::= SEQUENCE OF SchedulingInfo-NB-v1530
// SIZE (1..maxSI-Message-NB-r13)
type SchedulinginfolistNbV1530 struct {
	Value []SchedulinginfoNbV1530 `lb:1,ub:maxSIMessageNbR13`
}
