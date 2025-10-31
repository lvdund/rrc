package ies

import "rrc/utils"

// EUTRA-MBSFN-SubframeConfig-subframeAllocation1 ::= CHOICE
const (
	EutraMbsfnSubframeconfigSubframeallocation1ChoiceNothing = iota
	EutraMbsfnSubframeconfigSubframeallocation1ChoiceOneframe
	EutraMbsfnSubframeconfigSubframeallocation1ChoiceFourframes
)

type EutraMbsfnSubframeconfigSubframeallocation1 struct {
	Choice     uint64
	Oneframe   *utils.BITSTRING `lb:6,ub:6`
	Fourframes *utils.BITSTRING `lb:24,ub:24`
}
