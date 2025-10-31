package ies

// ResourceReservationConfig-NB-r16-resourceReservation-r16 ::= CHOICE
const (
	ResourcereservationconfigNbR16ResourcereservationR16ChoiceNothing = iota
	ResourcereservationconfigNbR16ResourcereservationR16ChoiceSubframebitmapR16
	ResourcereservationconfigNbR16ResourcereservationR16ChoiceSlotconfigR16
)

type ResourcereservationconfigNbR16ResourcereservationR16 struct {
	Choice            uint64
	SubframebitmapR16 *ResourcereservationconfigNbR16ResourcereservationR16SubframebitmapR16
	SlotconfigR16     *ResourcereservationconfigNbR16ResourcereservationR16SlotconfigR16
}
