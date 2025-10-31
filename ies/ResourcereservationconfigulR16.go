package ies

import "rrc/utils"

// ResourceReservationConfigUL-r16 ::= SEQUENCE
// Extensible
type ResourcereservationconfigulR16 struct {
	PeriodicitystartposR16 PeriodicitystartposR16
	SlotbitmapR16          *ResourcereservationconfigulR16SlotbitmapR16
	Symbolbitmap1R16       *utils.BITSTRING `lb:7,ub:7`
	Symbolbitmap2R16       *utils.BITSTRING `lb:7,ub:7`
}
