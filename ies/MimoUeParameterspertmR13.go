package ies

// MIMO-UE-ParametersPerTM-r13 ::= SEQUENCE
type MimoUeParameterspertmR13 struct {
	NonprecodedR13            *MimoNonprecodedcapabilitiesR13
	BeamformedR13             *MimoUeBeamformedcapabilitiesR13
	ChannelmeasrestrictionR13 *MimoUeParameterspertmR13ChannelmeasrestrictionR13
	DmrsEnhancementsR13       *MimoUeParameterspertmR13DmrsEnhancementsR13
	CsiRsEnhancementstddR13   *MimoUeParameterspertmR13CsiRsEnhancementstddR13
}
