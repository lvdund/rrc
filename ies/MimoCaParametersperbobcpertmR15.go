package ies

// MIMO-CA-ParametersPerBoBCPerTM-r15 ::= SEQUENCE
type MimoCaParametersperbobcpertmR15 struct {
	NonprecodedR13          *MimoNonprecodedcapabilitiesR13
	BeamformedR13           *MimoBeamformedcapabilitylistR13
	DmrsEnhancementsR13     *MimoCaParametersperbobcpertmR15DmrsEnhancementsR13
	CsiReportingnpR14       *MimoCaParametersperbobcpertmR15CsiReportingnpR14
	CsiReportingadvancedR14 *MimoCaParametersperbobcpertmR15CsiReportingadvancedR14
}
