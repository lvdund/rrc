package ies

import "rrc/utils"

// RLF-TimersAndConstants-n311 ::= ENUMERATED
type RlfTimersandconstantsN311 struct {
	Value utils.ENUMERATED
}

const (
	RlfTimersandconstantsN311EnumeratedNothing = iota
	RlfTimersandconstantsN311EnumeratedN1
	RlfTimersandconstantsN311EnumeratedN2
	RlfTimersandconstantsN311EnumeratedN3
	RlfTimersandconstantsN311EnumeratedN4
	RlfTimersandconstantsN311EnumeratedN5
	RlfTimersandconstantsN311EnumeratedN6
	RlfTimersandconstantsN311EnumeratedN8
	RlfTimersandconstantsN311EnumeratedN10
)
