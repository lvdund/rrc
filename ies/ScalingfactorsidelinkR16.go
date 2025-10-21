package ies

import "rrc/utils"

// ScalingFactorSidelink-r16 ::= ENUMERATED
type ScalingfactorsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	ScalingfactorsidelinkR16F0p4  = 0
	ScalingfactorsidelinkR16F0p75 = 1
	ScalingfactorsidelinkR16F0p8  = 2
	ScalingfactorsidelinkR16F1    = 3
)
