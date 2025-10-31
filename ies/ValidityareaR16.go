package ies

// ValidityArea-r16 ::= SEQUENCE
type ValidityareaR16 struct {
	CarrierfreqR16      ArfcnValueeutraR9
	ValiditycelllistR16 *ValiditycelllistR16
}
