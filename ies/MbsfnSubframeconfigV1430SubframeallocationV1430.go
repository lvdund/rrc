package ies

import "rrc/utils"

// MBSFN-SubframeConfig-v1430-subframeAllocation-v1430 ::= CHOICE
const (
	MbsfnSubframeconfigV1430SubframeallocationV1430ChoiceNothing = iota
	MbsfnSubframeconfigV1430SubframeallocationV1430ChoiceOneframeV1430
	MbsfnSubframeconfigV1430SubframeallocationV1430ChoiceFourframesV1430
)

type MbsfnSubframeconfigV1430SubframeallocationV1430 struct {
	Choice          uint64
	OneframeV1430   *utils.BITSTRING `lb:2,ub:2`
	FourframesV1430 *utils.BITSTRING `lb:8,ub:8`
}
