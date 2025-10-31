package ies

import "rrc/utils"

// ResourceReservationConfig-NB-r16-resourceReservation-r16-slotConfig-r16-symbolBitmap-r16-symbolBitmapFddUlOrTdd ::= SEQUENCE
type ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16SymbolbitmapR16Symbolbitmapfddulortdd struct {
	Symbolbitmap1R16 *utils.BITSTRING `lb:7,ub:7`
	Symbolbitmap2R16 *utils.BITSTRING `lb:7,ub:7`
}
