package ies

import "rrc/utils"

// ResourceReservationConfigDL-r16-slotBitmap-r16 ::= CHOICE
const (
	ResourcereservationconfigdlR16SlotbitmapR16ChoiceNothing = iota
	ResourcereservationconfigdlR16SlotbitmapR16ChoiceSlotpattern10ms
	ResourcereservationconfigdlR16SlotbitmapR16ChoiceSlotpattern40ms
)

type ResourcereservationconfigdlR16SlotbitmapR16 struct {
	Choice          uint64
	Slotpattern10ms *utils.BITSTRING `lb:20,ub:20`
	Slotpattern40ms *utils.BITSTRING `lb:80,ub:80`
}
