package ies

import "rrc/utils"

// FeatureSetDL-PerCC-r15 ::= SEQUENCE
type FeaturesetdlPerccR15 struct {
	Fourlayertm3Tm4R15               *utils.ENUMERATED
	SupportedmimoCapabilitydlMrdcR15 *MimoCapabilitydlR10
	SupportedcsiProcR15              *utils.ENUMERATED
}
