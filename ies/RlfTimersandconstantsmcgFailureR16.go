package ies

import "rrc/utils"

// RLF-TimersAndConstantsMCG-Failure-r16 ::= CHOICE
// Extensible
type RlfTimersandconstantsmcgFailureR16 interface {
	isRlfTimersandconstantsmcgFailureR16()
}

type RlfTimersandconstantsmcgFailureR16Release struct {
	Value struct{}
}

func (*RlfTimersandconstantsmcgFailureR16Release) isRlfTimersandconstantsmcgFailureR16() {}

type RlfTimersandconstantsmcgFailureR16Setup struct {
	Value interface{}
}

func (*RlfTimersandconstantsmcgFailureR16Setup) isRlfTimersandconstantsmcgFailureR16() {}
