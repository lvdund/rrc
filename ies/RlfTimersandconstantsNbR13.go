package ies

import "rrc/utils"

// RLF-TimersAndConstants-NB-r13 ::= CHOICE
// Extensible
type RlfTimersandconstantsNbR13 interface {
	isRlfTimersandconstantsNbR13()
}

type RlfTimersandconstantsNbR13Release struct {
	Value struct{}
}

func (*RlfTimersandconstantsNbR13Release) isRlfTimersandconstantsNbR13() {}

type RlfTimersandconstantsNbR13Setup struct {
	Value interface{}
}

func (*RlfTimersandconstantsNbR13Setup) isRlfTimersandconstantsNbR13() {}
