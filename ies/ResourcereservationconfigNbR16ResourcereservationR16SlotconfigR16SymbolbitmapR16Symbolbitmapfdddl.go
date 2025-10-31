package ies

import "rrc/utils"

// ResourceReservationConfig-NB-r16-resourceReservation-r16-slotConfig-r16-symbolBitmap-r16-symbolBitmapFddDl ::= SEQUENCE
type ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16SymbolbitmapR16Symbolbitmapfdddl struct {
	Symbolbitmap1R16 *utils.BITSTRING `lb:5,ub:5`
	Symbolbitmap2R16 *utils.BITSTRING `lb:5,ub:5`
}
