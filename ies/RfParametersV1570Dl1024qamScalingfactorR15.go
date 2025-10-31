package ies

import "rrc/utils"

// RF-Parameters-v1570-dl-1024QAM-ScalingFactor-r15 ::= ENUMERATED
type RfParametersV1570Dl1024qamScalingfactorR15 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1570Dl1024qamScalingfactorR15EnumeratedNothing = iota
	RfParametersV1570Dl1024qamScalingfactorR15EnumeratedV1
	RfParametersV1570Dl1024qamScalingfactorR15EnumeratedV1dot2
	RfParametersV1570Dl1024qamScalingfactorR15EnumeratedV1dot25
)
