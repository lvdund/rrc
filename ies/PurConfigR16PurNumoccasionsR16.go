package ies

import "rrc/utils"

// PUR-Config-r16-pur-NumOccasions-r16 ::= ENUMERATED
type PurConfigR16PurNumoccasionsR16 struct {
	Value utils.ENUMERATED
}

const (
	PurConfigR16PurNumoccasionsR16EnumeratedNothing = iota
	PurConfigR16PurNumoccasionsR16EnumeratedOne
	PurConfigR16PurNumoccasionsR16EnumeratedInfinite
)
