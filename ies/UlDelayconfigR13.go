package ies

import "rrc/utils"

// UL-DelayConfig-r13 ::= CHOICE
type UlDelayconfigR13 interface {
	isUlDelayconfigR13()
}

type UlDelayconfigR13Release struct {
	Value struct{}
}

func (*UlDelayconfigR13Release) isUlDelayconfigR13() {}

type UlDelayconfigR13Setup struct {
	Value interface{}
}

func (*UlDelayconfigR13Setup) isUlDelayconfigR13() {}
