package ies

import "rrc/utils"

// FeatureSetDL-PerCC-r15-fourLayerTM3-TM4-r15 ::= ENUMERATED
type FeaturesetdlPerccR15Fourlayertm3Tm4R15 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdlPerccR15Fourlayertm3Tm4R15EnumeratedNothing = iota
	FeaturesetdlPerccR15Fourlayertm3Tm4R15EnumeratedSupported
)
