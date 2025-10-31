package ies

import "rrc/utils"

// UE-TimersAndConstants-n311 ::= ENUMERATED
type UeTimersandconstantsN311 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsN311EnumeratedNothing = iota
	UeTimersandconstantsN311EnumeratedN1
	UeTimersandconstantsN311EnumeratedN2
	UeTimersandconstantsN311EnumeratedN3
	UeTimersandconstantsN311EnumeratedN4
	UeTimersandconstantsN311EnumeratedN5
	UeTimersandconstantsN311EnumeratedN6
	UeTimersandconstantsN311EnumeratedN8
	UeTimersandconstantsN311EnumeratedN10
)
