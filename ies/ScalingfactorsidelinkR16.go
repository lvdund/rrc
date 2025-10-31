package ies

import "rrc/utils"

// ScalingFactorSidelink-r16 ::= ENUMERATED
type ScalingfactorsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	ScalingfactorsidelinkR16EnumeratedNothing = iota
	ScalingfactorsidelinkR16EnumeratedF0p4
	ScalingfactorsidelinkR16EnumeratedF0p75
	ScalingfactorsidelinkR16EnumeratedF0p8
	ScalingfactorsidelinkR16EnumeratedF1
)
