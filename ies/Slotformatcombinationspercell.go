package ies

import "rrc/utils"

// SlotFormatCombinationsPerCell ::= SEQUENCE
// Extensible
type Slotformatcombinationspercell struct {
	Servingcellid          Servcellindex
	Subcarrierspacing      Subcarrierspacing
	Subcarrierspacing2     *Subcarrierspacing
	Slotformatcombinations *[]Slotformatcombination                            `lb:1,ub:maxNrofSlotFormatCombinationsPerSet`
	Positionindci          *utils.INTEGER                                      `lb:0,ub:maxSFIDciPayloadsize1`
	EnableconfiguredulR16  *SlotformatcombinationspercellEnableconfiguredulR16 `ext`
}
