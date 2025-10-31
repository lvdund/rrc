package ies

// MIMO-UE-ParametersPerTM-v1430 ::= SEQUENCE
type MimoUeParameterspertmV1430 struct {
	NzpCsiRsAperiodicinfoR14 *MimoUeParameterspertmV1430NzpCsiRsAperiodicinfoR14
	NzpCsiRsPeriodicinfoR14  *MimoUeParameterspertmV1430NzpCsiRsPeriodicinfoR14
	ZpCsiRsAperiodicinfoR14  *MimoUeParameterspertmV1430ZpCsiRsAperiodicinfoR14
	UlDmrsEnhancementsR14    *MimoUeParameterspertmV1430UlDmrsEnhancementsR14
	DensityreductionnpR14    *MimoUeParameterspertmV1430DensityreductionnpR14
	DensityreductionbfR14    *MimoUeParameterspertmV1430DensityreductionbfR14
	HybridcsiR14             *MimoUeParameterspertmV1430HybridcsiR14
	SemiolR14                *MimoUeParameterspertmV1430SemiolR14
	CsiReportingnpR14        *MimoUeParameterspertmV1430CsiReportingnpR14
	CsiReportingadvancedR14  *MimoUeParameterspertmV1430CsiReportingadvancedR14
}
