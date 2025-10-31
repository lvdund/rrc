package ies

import "rrc/utils"

// Phy-ParametersCommon-bwp-SwitchingMultiDormancyCCs-r16 ::= CHOICE
const (
	PhyParameterscommonBwpSwitchingmultidormancyccsR16ChoiceNothing = iota
	PhyParameterscommonBwpSwitchingmultidormancyccsR16ChoiceType1R16
	PhyParameterscommonBwpSwitchingmultidormancyccsR16ChoiceType2R16
)

type PhyParameterscommonBwpSwitchingmultidormancyccsR16 struct {
	Choice   uint64
	Type1R16 *utils.ENUMERATED
	Type2R16 *utils.ENUMERATED
}
