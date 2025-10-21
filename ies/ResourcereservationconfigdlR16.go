package ies

import "rrc/utils"

// ResourceReservationConfigDL-r16 ::= SEQUENCE
// Extensible
type ResourcereservationconfigdlR16 struct {
	PeriodicitystartposR16     PeriodicitystartposR16
	ResourcereservationfreqR16 *interface{}
	SlotbitmapR16              interface{}
	Symbolbitmap1R16           *utils.BITSTRING
	Symbolbitmap2R16           *utils.BITSTRING
}
