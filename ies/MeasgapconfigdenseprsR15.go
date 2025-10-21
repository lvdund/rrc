package ies

import "rrc/utils"

// MeasGapConfigDensePRS-r15 ::= CHOICE
// Extensible
type MeasgapconfigdenseprsR15 interface {
	isMeasgapconfigdenseprsR15()
}

type MeasgapconfigdenseprsR15Release struct {
	Value struct{}
}

func (*MeasgapconfigdenseprsR15Release) isMeasgapconfigdenseprsR15() {}

type MeasgapconfigdenseprsR15Setup struct {
	Value interface{}
}

func (*MeasgapconfigdenseprsR15Setup) isMeasgapconfigdenseprsR15() {}
