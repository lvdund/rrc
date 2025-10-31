package ies

// ValidityArea-r16 ::= SEQUENCE
type ValidityareaR16 struct {
	CarrierfreqR16      ArfcnValuenr
	ValiditycelllistR16 *Validitycelllist
}
