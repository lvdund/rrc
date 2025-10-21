package ies

import "rrc/utils"

// HighSpeedEnhParameters-r14 ::= SEQUENCE
type HighspeedenhparametersR14 struct {
	MeasurementenhancementsR14  *utils.ENUMERATED
	DemodulationenhancementsR14 *utils.ENUMERATED
	PrachEnhancementsR14        *utils.ENUMERATED
}
