package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15-setup-tbs-scalingFactorSubslotSPS-UL-Repetitions-r15 ::= ENUMERATED
type SpsConfigulSttiR15SetupTbsScalingfactorsubslotspsUlRepetitionsR15 struct {
	Value utils.ENUMERATED
}

const (
	SpsConfigulSttiR15SetupTbsScalingfactorsubslotspsUlRepetitionsR15EnumeratedNothing = iota
	SpsConfigulSttiR15SetupTbsScalingfactorsubslotspsUlRepetitionsR15EnumeratedN6
	SpsConfigulSttiR15SetupTbsScalingfactorsubslotspsUlRepetitionsR15EnumeratedN12
)
