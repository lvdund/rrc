package ies

import "rrc/utils"

// LogicalChGroupInfoList-r13 ::= SEQUENCE OF SL-PriorityList-r13
// SIZE (1..maxLCG-r13)
type LogicalchgroupinfolistR13 struct {
	Value []SlPrioritylistR13
}
