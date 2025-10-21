package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO-v1480 ::= CHOICE
type CsiRsConfigemimoV1480 interface {
	isCsiRsConfigemimoV1480()
}

type CsiRsConfigemimoV1480Release struct {
	Value struct{}
}

func (*CsiRsConfigemimoV1480Release) isCsiRsConfigemimoV1480() {}

type CsiRsConfigemimoV1480Setup struct {
	Value interface{}
}

func (*CsiRsConfigemimoV1480Setup) isCsiRsConfigemimoV1480() {}
