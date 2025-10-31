package ies

// SL-RLC-Config-r16 ::= CHOICE
// Extensible
const (
	SlRlcConfigR16ChoiceNothing = iota
	SlRlcConfigR16ChoiceSlAmRlcR16
	SlRlcConfigR16ChoiceSlUmRlcR16
)

type SlRlcConfigR16 struct {
	Choice     uint64
	SlAmRlcR16 *SlRlcConfigR16SlAmRlcR16
	SlUmRlcR16 *SlRlcConfigR16SlUmRlcR16
}
