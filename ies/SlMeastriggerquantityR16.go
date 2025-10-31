package ies

// SL-MeasTriggerQuantity-r16 ::= CHOICE
// Extensible
const (
	SlMeastriggerquantityR16ChoiceNothing = iota
	SlMeastriggerquantityR16ChoiceSlRsrpR16
)

type SlMeastriggerquantityR16 struct {
	Choice    uint64
	SlRsrpR16 *RsrpRange
}
