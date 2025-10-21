package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-r13 ::= SEQUENCE
type MimoUeParameterspertmR13 struct {
	NonprecodedR13            *MimoNonprecodedcapabilitiesR13
	BeamformedR13             *MimoUeBeamformedcapabilitiesR13
	ChannelmeasrestrictionR13 *utils.ENUMERATED
	DmrsEnhancementsR13       *utils.ENUMERATED
	CsiRsEnhancementstddR13   *utils.ENUMERATED
}
