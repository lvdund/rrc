package ies

import "rrc/utils"

// ResourceReservationConfig-NB-r16-resourceReservation-r16-subframeBitmap-r16 ::= CHOICE
const (
	ResourcereservationconfigNbR16ResourcereservationR16SubframebitmapR16ChoiceNothing = iota
	ResourcereservationconfigNbR16ResourcereservationR16SubframebitmapR16ChoiceSubframepattern10ms
	ResourcereservationconfigNbR16ResourcereservationR16SubframebitmapR16ChoiceSubframepattern40ms
)

type ResourcereservationconfigNbR16ResourcereservationR16SubframebitmapR16 struct {
	Choice              uint64
	Subframepattern10ms *utils.BITSTRING `lb:10,ub:10`
	Subframepattern40ms *utils.BITSTRING `lb:40,ub:40`
}
