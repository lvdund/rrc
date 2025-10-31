package ies

import "rrc/utils"

// PUCCH-Config-subslotLengthForPUCCH-r16 ::= CHOICE
const (
	PucchConfigSubslotlengthforpucchR16ChoiceNothing = iota
	PucchConfigSubslotlengthforpucchR16ChoiceNormalcpR16
	PucchConfigSubslotlengthforpucchR16ChoiceExtendedcpR16
)

type PucchConfigSubslotlengthforpucchR16 struct {
	Choice        uint64
	NormalcpR16   *utils.ENUMERATED
	ExtendedcpR16 *utils.ENUMERATED
}
