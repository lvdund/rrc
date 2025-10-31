package ies

import "rrc/utils"

// Phy-ParametersCommon-bwp-SwitchingMultiCCs-r16 ::= CHOICE
const (
	PhyParameterscommonBwpSwitchingmulticcsR16ChoiceNothing = iota
	PhyParameterscommonBwpSwitchingmulticcsR16ChoiceType1R16
	PhyParameterscommonBwpSwitchingmulticcsR16ChoiceType2R16
)

type PhyParameterscommonBwpSwitchingmulticcsR16 struct {
	Choice   uint64
	Type1R16 *utils.ENUMERATED
	Type2R16 *utils.ENUMERATED
}
