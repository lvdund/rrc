package ies

// SL-RLC-ConfigPC5-r16 ::= CHOICE
// Extensible
const (
	SlRlcConfigpc5R16ChoiceNothing = iota
	SlRlcConfigpc5R16ChoiceSlAmRlcR16
	SlRlcConfigpc5R16ChoiceSlUmBiDirectionalRlcR16
	SlRlcConfigpc5R16ChoiceSlUmUniDirectionalRlcR16
)

type SlRlcConfigpc5R16 struct {
	Choice                   uint64
	SlAmRlcR16               *SlRlcConfigpc5R16SlAmRlcR16
	SlUmBiDirectionalRlcR16  *SlRlcConfigpc5R16SlUmBiDirectionalRlcR16
	SlUmUniDirectionalRlcR16 *SlRlcConfigpc5R16SlUmUniDirectionalRlcR16
}
