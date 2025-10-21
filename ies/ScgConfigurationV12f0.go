package ies

import "rrc/utils"

// SCG-Configuration-v12f0 ::= CHOICE
type ScgConfigurationV12f0 interface {
	isScgConfigurationV12f0()
}

type ScgConfigurationV12f0Release struct {
	Value struct{}
}

func (*ScgConfigurationV12f0Release) isScgConfigurationV12f0() {}

type ScgConfigurationV12f0Setup struct {
	Value interface{}
}

func (*ScgConfigurationV12f0Setup) isScgConfigurationV12f0() {}
