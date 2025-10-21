package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430 ::= SEQUENCE
type MimoUeParameterspertmV1430 struct {
	NzpCsiRsAperiodicinfoR14 *interface{}
	NzpCsiRsPeriodicinfoR14  *interface{}
	ZpCsiRsAperiodicinfoR14  *utils.ENUMERATED
	UlDmrsEnhancementsR14    *utils.ENUMERATED
	DensityreductionnpR14    *utils.ENUMERATED
	DensityreductionbfR14    *utils.ENUMERATED
	HybridcsiR14             *utils.ENUMERATED
	SemiolR14                *utils.ENUMERATED
	CsiReportingnpR14        *utils.ENUMERATED
	CsiReportingadvancedR14  *utils.ENUMERATED
}
