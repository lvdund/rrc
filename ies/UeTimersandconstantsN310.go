package ies

import "rrc/utils"

// UE-TimersAndConstants-n310 ::= ENUMERATED
type UeTimersandconstantsN310 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsN310EnumeratedNothing = iota
	UeTimersandconstantsN310EnumeratedN1
	UeTimersandconstantsN310EnumeratedN2
	UeTimersandconstantsN310EnumeratedN3
	UeTimersandconstantsN310EnumeratedN4
	UeTimersandconstantsN310EnumeratedN6
	UeTimersandconstantsN310EnumeratedN8
	UeTimersandconstantsN310EnumeratedN10
	UeTimersandconstantsN310EnumeratedN20
)
