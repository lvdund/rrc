package ies

import "rrc/utils"

// SSB-ToMeasure-r15 ::= CHOICE
type SsbTomeasureR15 interface {
	isSsbTomeasureR15()
}

type SsbTomeasureR15ShortbitmapR15 struct {
	Value utils.BITSTRING
}

func (*SsbTomeasureR15ShortbitmapR15) isSsbTomeasureR15() {}

type SsbTomeasureR15MediumbitmapR15 struct {
	Value utils.BITSTRING
}

func (*SsbTomeasureR15MediumbitmapR15) isSsbTomeasureR15() {}

type SsbTomeasureR15LongbitmapR15 struct {
	Value utils.BITSTRING
}

func (*SsbTomeasureR15LongbitmapR15) isSsbTomeasureR15() {}
