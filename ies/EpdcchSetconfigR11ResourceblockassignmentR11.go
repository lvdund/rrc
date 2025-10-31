package ies

import "rrc/utils"

// EPDCCH-SetConfig-r11-resourceBlockAssignment-r11 ::= SEQUENCE
type EpdcchSetconfigR11ResourceblockassignmentR11 struct {
	NumberprbPairsR11          EpdcchSetconfigR11ResourceblockassignmentR11NumberprbPairsR11
	ResourceblockassignmentR11 utils.BITSTRING `lb:4,ub:38`
}
