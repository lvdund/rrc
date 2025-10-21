package ies

import "rrc/utils"

// LogicalChGroupInfoList-v1530 ::= SEQUENCE OF SL-ReliabilityList-r15
// SIZE (1..maxLCG-r13)
type LogicalchgroupinfolistV1530 struct {
	Value []SlReliabilitylistR15
}
