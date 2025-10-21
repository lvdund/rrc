package ies

import "rrc/utils"

// SkipSubframeProcessing-r15 ::= SEQUENCE
type SkipsubframeprocessingR15 struct {
	SkipprocessingdlSlotR15    *utils.INTEGER
	SkipprocessingdlSubslotR15 *utils.INTEGER
	SkipprocessingulSlotR15    *utils.INTEGER
	SkipprocessingulSubslotR15 *utils.INTEGER
}
