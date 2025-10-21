package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO-v1430 ::= CHOICE
type CsiRsConfigemimoV1430 interface {
	isCsiRsConfigemimoV1430()
}

type CsiRsConfigemimoV1430Release struct {
	Value struct{}
}

func (*CsiRsConfigemimoV1430Release) isCsiRsConfigemimoV1430() {}

type CsiRsConfigemimoV1430Setup struct {
	Value interface{}
}

func (*CsiRsConfigemimoV1430Setup) isCsiRsConfigemimoV1430() {}
