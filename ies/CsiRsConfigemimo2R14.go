package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO2-r14 ::= CHOICE
type CsiRsConfigemimo2R14 interface {
	isCsiRsConfigemimo2R14()
}

type CsiRsConfigemimo2R14Release struct {
	Value struct{}
}

func (*CsiRsConfigemimo2R14Release) isCsiRsConfigemimo2R14() {}

type CsiRsConfigemimo2R14Setup struct {
	Value CsiRsConfigbeamformedR14
}

func (*CsiRsConfigemimo2R14Setup) isCsiRsConfigemimo2R14() {}
