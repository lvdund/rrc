package ies

import "rrc/utils"

// UL-DelayValueConfig-r16 ::= CHOICE
type UlDelayvalueconfigR16 interface {
	isUlDelayvalueconfigR16()
}

type UlDelayvalueconfigR16Release struct {
	Value struct{}
}

func (*UlDelayvalueconfigR16Release) isUlDelayvalueconfigR16() {}

type UlDelayvalueconfigR16Setup struct {
	Value interface{}
}

func (*UlDelayvalueconfigR16Setup) isUlDelayvalueconfigR16() {}
