package ies

import "rrc/utils"

// ServingCellConfig-ca-SlotOffset-r16 ::= CHOICE
const (
	ServingcellconfigCaSlotoffsetR16ChoiceNothing = iota
	ServingcellconfigCaSlotoffsetR16ChoiceRefscs15khz
	ServingcellconfigCaSlotoffsetR16ChoiceRefscs30khz
	ServingcellconfigCaSlotoffsetR16ChoiceRefscs60khz
	ServingcellconfigCaSlotoffsetR16ChoiceRefscs120khz
)

type ServingcellconfigCaSlotoffsetR16 struct {
	Choice       uint64
	Refscs15khz  *utils.INTEGER `lb:-2,ub:2`
	Refscs30khz  *utils.INTEGER `lb:-5,ub:5`
	Refscs60khz  *utils.INTEGER `lb:-10,ub:10`
	Refscs120khz *utils.INTEGER `lb:-20,ub:20`
}
