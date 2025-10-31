package ies

import "rrc/utils"

// ResourceReservationConfigUL-r16-slotBitmap-r16 ::= CHOICE
const (
	ResourcereservationconfigulR16SlotbitmapR16ChoiceNothing = iota
	ResourcereservationconfigulR16SlotbitmapR16ChoiceSlotpattern10ms
	ResourcereservationconfigulR16SlotbitmapR16ChoiceSlotpattern40ms
)

type ResourcereservationconfigulR16SlotbitmapR16 struct {
	Choice          uint64
	Slotpattern10ms *utils.BITSTRING `lb:20,ub:20`
	Slotpattern40ms *utils.BITSTRING `lb:80,ub:80`
}
