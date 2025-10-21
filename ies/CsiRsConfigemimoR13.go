package ies

import "rrc/utils"

// CSI-RS-ConfigEMIMO-r13 ::= CHOICE
type CsiRsConfigemimoR13 interface {
	isCsiRsConfigemimoR13()
}

type CsiRsConfigemimoR13Release struct {
	Value struct{}
}

func (*CsiRsConfigemimoR13Release) isCsiRsConfigemimoR13() {}

type CsiRsConfigemimoR13Setup struct {
	Value interface{}
}

func (*CsiRsConfigemimoR13Setup) isCsiRsConfigemimoR13() {}
