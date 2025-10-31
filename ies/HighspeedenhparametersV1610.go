package ies

// HighSpeedEnhParameters-v1610 ::= SEQUENCE
type HighspeedenhparametersV1610 struct {
	MeasurementenhancementsscellR16 *HighspeedenhparametersV1610MeasurementenhancementsscellR16
	Measurementenhancements2R16     *HighspeedenhparametersV1610Measurementenhancements2R16
	Demodulationenhancements2R16    *HighspeedenhparametersV1610Demodulationenhancements2R16
	InterratEnhancementnrR16        *HighspeedenhparametersV1610InterratEnhancementnrR16
}
