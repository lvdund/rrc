package ies

import "rrc/utils"

// EIMTA-MainConfigServCell-r12 ::= CHOICE
type EimtaMainconfigservcellR12 interface {
	isEimtaMainconfigservcellR12()
}

type EimtaMainconfigservcellR12Release struct {
	Value struct{}
}

func (*EimtaMainconfigservcellR12Release) isEimtaMainconfigservcellR12() {}

type EimtaMainconfigservcellR12Setup struct {
	Value interface{}
}

func (*EimtaMainconfigservcellR12Setup) isEimtaMainconfigservcellR12() {}
