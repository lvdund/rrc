package ies

// HighSpeedEnhParameters-r14 ::= SEQUENCE
type HighspeedenhparametersR14 struct {
	MeasurementenhancementsR14  *HighspeedenhparametersR14MeasurementenhancementsR14
	DemodulationenhancementsR14 *HighspeedenhparametersR14DemodulationenhancementsR14
	PrachEnhancementsR14        *HighspeedenhparametersR14PrachEnhancementsR14
}
