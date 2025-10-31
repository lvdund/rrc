package ies

// LogicalChGroupInfoList-v1530 ::= SEQUENCE OF SL-ReliabilityList-r15
// SIZE (1..maxLCG-r13)
type LogicalchgroupinfolistV1530 struct {
	Value []SlReliabilitylistR15 `lb:1,ub:maxLCGR13`
}
