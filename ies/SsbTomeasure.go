package ies

import "rrc/utils"

// SSB-ToMeasure ::= CHOICE
const (
	SsbTomeasureChoiceNothing = iota
	SsbTomeasureChoiceShortbitmap
	SsbTomeasureChoiceMediumbitmap
	SsbTomeasureChoiceLongbitmap
)

type SsbTomeasure struct {
	Choice       uint64
	Shortbitmap  *utils.BITSTRING `lb:4,ub:4`
	Mediumbitmap *utils.BITSTRING `lb:8,ub:8`
	Longbitmap   *utils.BITSTRING `lb:64,ub:64`
}
