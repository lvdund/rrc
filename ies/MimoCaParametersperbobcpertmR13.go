package ies

// MIMO-CA-ParametersPerBoBCPerTM-r13 ::= SEQUENCE
type MimoCaParametersperbobcpertmR13 struct {
	NonprecodedR13      *MimoNonprecodedcapabilitiesR13
	BeamformedR13       *MimoBeamformedcapabilitylistR13
	DmrsEnhancementsR13 *MimoCaParametersperbobcpertmR13DmrsEnhancementsR13
}
