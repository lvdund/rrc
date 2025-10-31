package ies

import "rrc/utils"

// MBSFN-SubframeConfig-v1610-subframeAllocation-v1610 ::= CHOICE
const (
	MbsfnSubframeconfigV1610SubframeallocationV1610ChoiceNothing = iota
	MbsfnSubframeconfigV1610SubframeallocationV1610ChoiceOneframeV1610
	MbsfnSubframeconfigV1610SubframeallocationV1610ChoiceFourframesV1610
)

type MbsfnSubframeconfigV1610SubframeallocationV1610 struct {
	Choice          uint64
	OneframeV1610   *utils.BITSTRING `lb:2,ub:2`
	FourframesV1610 *utils.BITSTRING `lb:8,ub:8`
}
