package ies

import "rrc/utils"

// CFI-Config-r15 ::= SEQUENCE
type CfiConfigR15 struct {
	CfiSubframenonmbsfnR15    *utils.INTEGER
	CfiSlotsubslotnonmbsfnR15 *utils.INTEGER
	CfiSubframembsfnR15       *utils.INTEGER
	CfiSlotsubslotmbsfnR15    *utils.INTEGER
}
