package ies

import "rrc/utils"

// FeatureSetUplink-v1710-mTRP-PUSCH-RepetitionTypeA-r17 ::= ENUMERATED
type FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17EnumeratedNothing = iota
	FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17EnumeratedN1
	FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17EnumeratedN2
	FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17EnumeratedN3
	FeaturesetuplinkV1710MtrpPuschRepetitiontypeaR17EnumeratedN4
)
