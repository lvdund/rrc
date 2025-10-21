package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO-Hybrid-r14 ::= CHOICE
type CsiRsConfigemimoHybridR14 interface {
	isCsiRsConfigemimoHybridR14()
}

type CsiRsConfigemimoHybridR14Release struct {
	Value struct{}
}

func (*CsiRsConfigemimoHybridR14Release) isCsiRsConfigemimoHybridR14() {}

type CsiRsConfigemimoHybridR14Setup struct {
	Value interface{}
}

func (*CsiRsConfigemimoHybridR14Setup) isCsiRsConfigemimoHybridR14() {}
