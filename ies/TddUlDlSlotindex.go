package ies

import "rrc/utils"

// TDD-UL-DL-SlotIndex ::= utils.INTEGER (0..maxNrofSlots-1)
type TddUlDlSlotindex struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSlots1`
}
