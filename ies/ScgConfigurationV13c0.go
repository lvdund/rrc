package ies

import "rrc/utils"

// SCG-Configuration-v13c0 ::= CHOICE
type ScgConfigurationV13c0 interface {
	isScgConfigurationV13c0()
}

type ScgConfigurationV13c0Release struct {
	Value struct{}
}

func (*ScgConfigurationV13c0Release) isScgConfigurationV13c0() {}

type ScgConfigurationV13c0Setup struct {
	Value interface{}
}

func (*ScgConfigurationV13c0Setup) isScgConfigurationV13c0() {}
