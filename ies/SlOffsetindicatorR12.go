package ies

import "rrc/utils"

// SL-OffsetIndicator-r12 ::= CHOICE
const (
	SlOffsetindicatorR12ChoiceNothing = iota
	SlOffsetindicatorR12ChoiceSmallR12
	SlOffsetindicatorR12ChoiceLargeR12
)

type SlOffsetindicatorR12 struct {
	Choice   uint64
	SmallR12 *utils.INTEGER `lb:0,ub:319`
	LargeR12 *utils.INTEGER `lb:0,ub:10239`
}
