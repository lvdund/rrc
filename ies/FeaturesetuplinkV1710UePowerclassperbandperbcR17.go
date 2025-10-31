package ies

import "rrc/utils"

// FeatureSetUplink-v1710-ue-PowerClassPerBandPerBC-r17 ::= ENUMERATED
type FeaturesetuplinkV1710UePowerclassperbandperbcR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710UePowerclassperbandperbcR17EnumeratedNothing = iota
	FeaturesetuplinkV1710UePowerclassperbandperbcR17EnumeratedPc1dot5
	FeaturesetuplinkV1710UePowerclassperbandperbcR17EnumeratedPc2
	FeaturesetuplinkV1710UePowerclassperbandperbcR17EnumeratedPc3
)
