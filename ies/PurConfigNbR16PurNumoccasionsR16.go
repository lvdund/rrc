package ies

import "rrc/utils"

// PUR-Config-NB-r16-pur-NumOccasions-r16 ::= ENUMERATED
type PurConfigNbR16PurNumoccasionsR16 struct {
	Value utils.ENUMERATED
}

const (
	PurConfigNbR16PurNumoccasionsR16EnumeratedNothing = iota
	PurConfigNbR16PurNumoccasionsR16EnumeratedOne
	PurConfigNbR16PurNumoccasionsR16EnumeratedInfinite
)
