package ies

import "rrc/utils"

// RA-Prioritization-scalingFactorBI ::= ENUMERATED
type RaPrioritizationScalingfactorbi struct {
	Value utils.ENUMERATED
}

const (
	RaPrioritizationScalingfactorbiEnumeratedNothing = iota
	RaPrioritizationScalingfactorbiEnumeratedZero
	RaPrioritizationScalingfactorbiEnumeratedDot25
	RaPrioritizationScalingfactorbiEnumeratedDot5
	RaPrioritizationScalingfactorbiEnumeratedDot75
)
