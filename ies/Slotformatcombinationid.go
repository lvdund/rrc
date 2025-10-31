package ies

import "rrc/utils"

// SlotFormatCombinationId ::= utils.INTEGER (0..maxNrofSlotFormatCombinationsPerSet-1)
type Slotformatcombinationid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSlotFormatCombinationsPerSet1`
}
