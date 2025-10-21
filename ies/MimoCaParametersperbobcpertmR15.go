package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-r15 ::= SEQUENCE
type MimoCaParametersperbobcpertmR15 struct {
	NonprecodedR13          *MimoNonprecodedcapabilitiesR13
	BeamformedR13           *MimoBeamformedcapabilitylistR13
	DmrsEnhancementsR13     *utils.ENUMERATED
	CsiReportingnpR14       *utils.ENUMERATED
	CsiReportingadvancedR14 *utils.ENUMERATED
}
