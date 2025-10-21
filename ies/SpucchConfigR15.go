package ies

import "rrc/utils"

// SPUCCH-Config-r15 ::= CHOICE
type SpucchConfigR15 interface {
	isSpucchConfigR15()
}

type SpucchConfigR15Release struct {
	Value struct{}
}

func (*SpucchConfigR15Release) isSpucchConfigR15() {}

type SpucchConfigR15Setup struct {
	Value interface{}
}

func (*SpucchConfigR15Setup) isSpucchConfigR15() {}
