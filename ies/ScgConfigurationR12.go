package ies

import "rrc/utils"

// SCG-Configuration-r12 ::= CHOICE
// Extensible
type ScgConfigurationR12 interface {
	isScgConfigurationR12()
}

type ScgConfigurationR12Release struct {
	Value struct{}
}

func (*ScgConfigurationR12Release) isScgConfigurationR12() {}

type ScgConfigurationR12Setup struct {
	Value interface{}
}

func (*ScgConfigurationR12Setup) isScgConfigurationR12() {}
