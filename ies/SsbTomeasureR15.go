package ies

import "rrc/utils"

// SSB-ToMeasure-r15 ::= CHOICE
const (
	SsbTomeasureR15ChoiceNothing = iota
	SsbTomeasureR15ChoiceShortbitmapR15
	SsbTomeasureR15ChoiceMediumbitmapR15
	SsbTomeasureR15ChoiceLongbitmapR15
)

type SsbTomeasureR15 struct {
	Choice          uint64
	ShortbitmapR15  *utils.BITSTRING `lb:4,ub:4`
	MediumbitmapR15 *utils.BITSTRING `lb:8,ub:8`
	LongbitmapR15   *utils.BITSTRING `lb:64,ub:64`
}
