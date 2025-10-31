package ies

import "rrc/utils"

// RepetitionSchemeConfig-r16 ::= CHOICE
const (
	RepetitionschemeconfigR16ChoiceNothing = iota
	RepetitionschemeconfigR16ChoiceFdmTdmR16
	RepetitionschemeconfigR16ChoiceSlotbasedR16
)

type RepetitionschemeconfigR16 struct {
	Choice       uint64
	FdmTdmR16    *utils.Setuprelease[FdmTdmR16]
	SlotbasedR16 *utils.Setuprelease[SlotbasedR16]
}
