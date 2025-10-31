package ies

import "rrc/utils"

// SL-SSB-TimeAllocation-r16 ::= SEQUENCE
type SlSsbTimeallocationR16 struct {
	SlNumssbWithinperiodR16 *SlSsbTimeallocationR16SlNumssbWithinperiodR16
	SlTimeoffsetssbR16      *utils.INTEGER `lb:0,ub:1279`
	SlTimeintervalR16       *utils.INTEGER `lb:0,ub:639`
}
