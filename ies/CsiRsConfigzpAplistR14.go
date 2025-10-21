package ies

import "rrc/utils"

// CSI-RS-ConfigZP-ApList-r14 ::= CHOICE
type CsiRsConfigzpAplistR14 interface {
	isCsiRsConfigzpAplistR14()
}

type CsiRsConfigzpAplistR14Release struct {
	Value struct{}
}

func (*CsiRsConfigzpAplistR14Release) isCsiRsConfigzpAplistR14() {}

type CsiRsConfigzpAplistR14Setup struct {
	Value interface{}
}

func (*CsiRsConfigzpAplistR14Setup) isCsiRsConfigzpAplistR14() {}
