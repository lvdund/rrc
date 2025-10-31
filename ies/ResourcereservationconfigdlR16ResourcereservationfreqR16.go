package ies

import "rrc/utils"

// ResourceReservationConfigDL-r16-resourceReservationFreq-r16 ::= CHOICE
const (
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceNothing = iota
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceRbgBitmap1dot4
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceRbgBitmap3
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceRbgBitmap5
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceRbgBitmap10
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceRbgBitmap15
	ResourcereservationconfigdlR16ResourcereservationfreqR16ChoiceRbgBitmap20
)

type ResourcereservationconfigdlR16ResourcereservationfreqR16 struct {
	Choice         uint64
	RbgBitmap1dot4 *utils.BITSTRING `lb:6,ub:6`
	RbgBitmap3     *utils.BITSTRING `lb:8,ub:8`
	RbgBitmap5     *utils.BITSTRING `lb:13,ub:13`
	RbgBitmap10    *utils.BITSTRING `lb:17,ub:17`
	RbgBitmap15    *utils.BITSTRING `lb:19,ub:19`
	RbgBitmap20    *utils.BITSTRING `lb:25,ub:25`
}
