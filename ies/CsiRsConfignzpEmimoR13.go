package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-EMIMO-r13 ::= CHOICE
type CsiRsConfignzpEmimoR13 interface {
	isCsiRsConfignzpEmimoR13()
}

type CsiRsConfignzpEmimoR13Release struct {
	Value struct{}
}

func (*CsiRsConfignzpEmimoR13Release) isCsiRsConfignzpEmimoR13() {}

type CsiRsConfignzpEmimoR13Setup struct {
	Value interface{}
}

func (*CsiRsConfignzpEmimoR13Setup) isCsiRsConfignzpEmimoR13() {}
