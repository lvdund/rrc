package ies

import "rrc/utils"

// HighSpeedEnhParameters-v1610 ::= SEQUENCE
type HighspeedenhparametersV1610 struct {
	MeasurementenhancementsscellR16 *utils.ENUMERATED
	Measurementenhancements2R16     *utils.ENUMERATED
	Demodulationenhancements2R16    *utils.ENUMERATED
	InterratEnhancementnrR16        *utils.ENUMERATED
}
