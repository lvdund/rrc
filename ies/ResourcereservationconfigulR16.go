package ies

import "rrc/utils"

// ResourceReservationConfigUL-r16 ::= SEQUENCE
// Extensible
type ResourcereservationconfigulR16 struct {
	PeriodicitystartposR16 PeriodicitystartposR16
	SlotbitmapR16          *interface{}
	Symbolbitmap1R16       *utils.BITSTRING
	Symbolbitmap2R16       *utils.BITSTRING
}
