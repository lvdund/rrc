package ies

import "rrc/utils"

// MBSFN-SubframeConfig-subframeAllocation ::= CHOICE
const (
	MbsfnSubframeconfigSubframeallocationChoiceNothing = iota
	MbsfnSubframeconfigSubframeallocationChoiceOneframe
	MbsfnSubframeconfigSubframeallocationChoiceFourframes
)

type MbsfnSubframeconfigSubframeallocation struct {
	Choice     uint64
	Oneframe   *utils.BITSTRING `lb:6,ub:6`
	Fourframes *utils.BITSTRING `lb:24,ub:24`
}
