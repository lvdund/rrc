package ies

// SL-RLC-ModeIndication-r16-sl-Mode-r16 ::= CHOICE
const (
	SlRlcModeindicationR16SlModeR16ChoiceNothing = iota
	SlRlcModeindicationR16SlModeR16ChoiceSlAmModeR16
	SlRlcModeindicationR16SlModeR16ChoiceSlUmModeR16
)

type SlRlcModeindicationR16SlModeR16 struct {
	Choice      uint64
	SlAmModeR16 *struct{}
	SlUmModeR16 *struct{}
}
