package ies

import "rrc/utils"

// SlotFormatCombination ::= SEQUENCE
type Slotformatcombination struct {
	Slotformatcombinationid Slotformatcombinationid
	Slotformats             []utils.INTEGER `lb:1,ub:maxNrofSlotFormatsPerCombination`
}
