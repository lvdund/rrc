package ies

import "rrc/utils"

// FeatureSetUplink-v1710-mTRP-PUSCH-TypeA-CB-r17 ::= ENUMERATED
type FeaturesetuplinkV1710MtrpPuschTypeaCbR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710MtrpPuschTypeaCbR17EnumeratedNothing = iota
	FeaturesetuplinkV1710MtrpPuschTypeaCbR17EnumeratedN1
	FeaturesetuplinkV1710MtrpPuschTypeaCbR17EnumeratedN2
	FeaturesetuplinkV1710MtrpPuschTypeaCbR17EnumeratedN4
)
