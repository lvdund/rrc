package ies

import "rrc/utils"

// RLF-TimersAndConstants-r13 ::= CHOICE
// Extensible
type RlfTimersandconstantsR13 interface {
	isRlfTimersandconstantsR13()
}

type RlfTimersandconstantsR13Release struct {
	Value struct{}
}

func (*RlfTimersandconstantsR13Release) isRlfTimersandconstantsR13() {}

type RlfTimersandconstantsR13Setup struct {
	Value interface{}
}

func (*RlfTimersandconstantsR13Setup) isRlfTimersandconstantsR13() {}
