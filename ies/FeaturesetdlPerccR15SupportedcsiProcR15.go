package ies

import "rrc/utils"

// FeatureSetDL-PerCC-r15-supportedCSI-Proc-r15 ::= ENUMERATED
type FeaturesetdlPerccR15SupportedcsiProcR15 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdlPerccR15SupportedcsiProcR15EnumeratedNothing = iota
	FeaturesetdlPerccR15SupportedcsiProcR15EnumeratedN1
	FeaturesetdlPerccR15SupportedcsiProcR15EnumeratedN3
	FeaturesetdlPerccR15SupportedcsiProcR15EnumeratedN4
)
