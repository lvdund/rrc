package ies

// FeatureSetDL-PerCC-r15 ::= SEQUENCE
type FeaturesetdlPerccR15 struct {
	Fourlayertm3Tm4R15               *FeaturesetdlPerccR15Fourlayertm3Tm4R15
	SupportedmimoCapabilitydlMrdcR15 *MimoCapabilitydlR10
	SupportedcsiProcR15              *FeaturesetdlPerccR15SupportedcsiProcR15
}
