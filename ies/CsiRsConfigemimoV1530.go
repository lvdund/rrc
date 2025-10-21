package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO-v1530 ::= CHOICE
type CsiRsConfigemimoV1530 interface {
	isCsiRsConfigemimoV1530()
}

type CsiRsConfigemimoV1530Release struct {
	Value struct{}
}

func (*CsiRsConfigemimoV1530Release) isCsiRsConfigemimoV1530() {}

type CsiRsConfigemimoV1530Setup struct {
	Value interface{}
}

func (*CsiRsConfigemimoV1530Setup) isCsiRsConfigemimoV1530() {}
