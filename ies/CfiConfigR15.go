package ies

import "rrc/utils"

// CFI-Config-r15 ::= SEQUENCE
type CfiConfigR15 struct {
	CfiSubframenonmbsfnR15    *utils.INTEGER `lb:0,ub:4`
	CfiSlotsubslotnonmbsfnR15 *utils.INTEGER `lb:0,ub:3`
	CfiSubframembsfnR15       *utils.INTEGER `lb:0,ub:2`
	CfiSlotsubslotmbsfnR15    *utils.INTEGER `lb:0,ub:2`
}
