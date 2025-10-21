package ies

import "rrc/utils"

// SPDCCH-Config-r15 ::= CHOICE
type SpdcchConfigR15 interface {
	isSpdcchConfigR15()
}

type SpdcchConfigR15Release struct {
	Value struct{}
}

func (*SpdcchConfigR15Release) isSpdcchConfigR15() {}

type SpdcchConfigR15Setup struct {
	Value interface{}
}

func (*SpdcchConfigR15Setup) isSpdcchConfigR15() {}
