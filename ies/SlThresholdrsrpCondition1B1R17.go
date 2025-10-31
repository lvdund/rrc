package ies

import "rrc/utils"

// SL-ThresholdRSRP-Condition1-B-1-r17 ::= SEQUENCE
type SlThresholdrsrpCondition1B1R17 struct {
	SlPriorityR17                  utils.INTEGER `lb:0,ub:8`
	SlThresholdrsrpCondition1B1R17 utils.INTEGER `lb:0,ub:66`
}
