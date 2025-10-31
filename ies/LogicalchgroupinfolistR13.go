package ies

// LogicalChGroupInfoList-r13 ::= SEQUENCE OF SL-PriorityList-r13
// SIZE (1..maxLCG-r13)
type LogicalchgroupinfolistR13 struct {
	Value []SlPrioritylistR13 `lb:1,ub:maxLCGR13`
}
