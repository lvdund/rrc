package ies

import "rrc/utils"

// RLC-BearerConfig-r15 ::= CHOICE
type RlcBearerconfigR15 interface {
	isRlcBearerconfigR15()
}

type RlcBearerconfigR15Release struct {
	Value struct{}
}

func (*RlcBearerconfigR15Release) isRlcBearerconfigR15() {}

type RlcBearerconfigR15Setup struct {
	Value interface{}
}

func (*RlcBearerconfigR15Setup) isRlcBearerconfigR15() {}
