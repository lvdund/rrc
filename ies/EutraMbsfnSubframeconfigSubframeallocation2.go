package ies

import "rrc/utils"

// EUTRA-MBSFN-SubframeConfig-subframeAllocation2 ::= CHOICE
const (
	EutraMbsfnSubframeconfigSubframeallocation2ChoiceNothing = iota
	EutraMbsfnSubframeconfigSubframeallocation2ChoiceOneframe
	EutraMbsfnSubframeconfigSubframeallocation2ChoiceFourframes
)

type EutraMbsfnSubframeconfigSubframeallocation2 struct {
	Choice     uint64
	Oneframe   *utils.BITSTRING `lb:2,ub:2`
	Fourframes *utils.BITSTRING `lb:8,ub:8`
}
