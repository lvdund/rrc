package ies

import "rrc/utils"

// FeatureSetDL-v1550-dl-1024QAM-r15 ::= ENUMERATED
type FeaturesetdlV1550Dl1024qamR15 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdlV1550Dl1024qamR15EnumeratedNothing = iota
	FeaturesetdlV1550Dl1024qamR15EnumeratedSupported
)
