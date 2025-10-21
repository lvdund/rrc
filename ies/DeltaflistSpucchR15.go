package ies

import "rrc/utils"

// DeltaFList-SPUCCH-r15 ::= CHOICE
// Extensible
type DeltaflistSpucchR15 interface {
	isDeltaflistSpucchR15()
}

type DeltaflistSpucchR15Release struct {
	Value struct{}
}

func (*DeltaflistSpucchR15Release) isDeltaflistSpucchR15() {}

type DeltaflistSpucchR15Setup struct {
	Value interface{}
}

func (*DeltaflistSpucchR15Setup) isDeltaflistSpucchR15() {}
