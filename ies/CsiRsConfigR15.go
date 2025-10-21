package ies

import "rrc/utils"

// CSI-RS-Config-r15 ::= CHOICE
type CsiRsConfigR15 interface {
	isCsiRsConfigR15()
}

type CsiRsConfigR15Release struct {
	Value struct{}
}

func (*CsiRsConfigR15Release) isCsiRsConfigR15() {}

type CsiRsConfigR15Setup struct {
	Value interface{}
}

func (*CsiRsConfigR15Setup) isCsiRsConfigR15() {}
