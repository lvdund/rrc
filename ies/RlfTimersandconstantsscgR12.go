package ies

import "rrc/utils"

// RLF-TimersAndConstantsSCG-r12 ::= CHOICE
// Extensible
type RlfTimersandconstantsscgR12 interface {
	isRlfTimersandconstantsscgR12()
}

type RlfTimersandconstantsscgR12Release struct {
	Value struct{}
}

func (*RlfTimersandconstantsscgR12Release) isRlfTimersandconstantsscgR12() {}

type RlfTimersandconstantsscgR12Setup struct {
	Value interface{}
}

func (*RlfTimersandconstantsscgR12Setup) isRlfTimersandconstantsscgR12() {}
