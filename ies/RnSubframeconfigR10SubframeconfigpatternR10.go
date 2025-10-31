package ies

import "rrc/utils"

// RN-SubframeConfig-r10-subframeConfigPattern-r10 ::= CHOICE
const (
	RnSubframeconfigR10SubframeconfigpatternR10ChoiceNothing = iota
	RnSubframeconfigR10SubframeconfigpatternR10ChoiceSubframeconfigpatternfddR10
	RnSubframeconfigR10SubframeconfigpatternR10ChoiceSubframeconfigpatterntddR10
)

type RnSubframeconfigR10SubframeconfigpatternR10 struct {
	Choice                      uint64
	SubframeconfigpatternfddR10 *utils.BITSTRING `lb:8,ub:8`
	SubframeconfigpatterntddR10 *utils.INTEGER   `lb:0,ub:31`
}
