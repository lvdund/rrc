package ies

import "rrc/utils"

// ResourceReservationConfigDL-r16 ::= SEQUENCE
// Extensible
type ResourcereservationconfigdlR16 struct {
	PeriodicitystartposR16     PeriodicitystartposR16
	ResourcereservationfreqR16 *ResourcereservationconfigdlR16ResourcereservationfreqR16
	SlotbitmapR16              ResourcereservationconfigdlR16SlotbitmapR16
	Symbolbitmap1R16           *utils.BITSTRING `lb:7,ub:7`
	Symbolbitmap2R16           *utils.BITSTRING `lb:7,ub:7`
}
