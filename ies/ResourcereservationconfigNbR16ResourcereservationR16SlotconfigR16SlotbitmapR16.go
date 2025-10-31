package ies

import "rrc/utils"

// ResourceReservationConfig-NB-r16-resourceReservation-r16-slotConfig-r16-slotBitmap-r16 ::= CHOICE
const (
	ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16SlotbitmapR16ChoiceNothing = iota
	ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16SlotbitmapR16ChoiceSlotpattern10ms
	ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16SlotbitmapR16ChoiceSlotpattern40ms
)

type ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16SlotbitmapR16 struct {
	Choice          uint64
	Slotpattern10ms *utils.BITSTRING `lb:20,ub:20`
	Slotpattern40ms *utils.BITSTRING `lb:80,ub:80`
}
