package ies

import "rrc/utils"

// RLF-TimersAndConstants-n310 ::= ENUMERATED
type RlfTimersandconstantsN310 struct {
	Value utils.ENUMERATED
}

const (
	RlfTimersandconstantsN310EnumeratedNothing = iota
	RlfTimersandconstantsN310EnumeratedN1
	RlfTimersandconstantsN310EnumeratedN2
	RlfTimersandconstantsN310EnumeratedN3
	RlfTimersandconstantsN310EnumeratedN4
	RlfTimersandconstantsN310EnumeratedN6
	RlfTimersandconstantsN310EnumeratedN8
	RlfTimersandconstantsN310EnumeratedN10
	RlfTimersandconstantsN310EnumeratedN20
)
