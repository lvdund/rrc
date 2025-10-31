package ies

import "rrc/utils"

// SkipSubframeProcessing-r15 ::= SEQUENCE
type SkipsubframeprocessingR15 struct {
	SkipprocessingdlSlotR15    *utils.INTEGER `lb:0,ub:3`
	SkipprocessingdlSubslotR15 *utils.INTEGER `lb:0,ub:3`
	SkipprocessingulSlotR15    *utils.INTEGER `lb:0,ub:3`
	SkipprocessingulSubslotR15 *utils.INTEGER `lb:0,ub:3`
}
