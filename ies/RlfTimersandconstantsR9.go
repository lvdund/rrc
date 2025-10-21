package ies

import "rrc/utils"

// RLF-TimersAndConstants-r9 ::= CHOICE
// Extensible
type RlfTimersandconstantsR9 interface {
	isRlfTimersandconstantsR9()
}

type RlfTimersandconstantsR9Release struct {
	Value struct{}
}

func (*RlfTimersandconstantsR9Release) isRlfTimersandconstantsR9() {}

type RlfTimersandconstantsR9Setup struct {
	Value interface{}
}

func (*RlfTimersandconstantsR9Setup) isRlfTimersandconstantsR9() {}
