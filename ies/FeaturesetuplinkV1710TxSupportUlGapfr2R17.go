package ies

import "rrc/utils"

// FeatureSetUplink-v1710-tx-Support-UL-GapFR2-r17 ::= ENUMERATED
type FeaturesetuplinkV1710TxSupportUlGapfr2R17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710TxSupportUlGapfr2R17EnumeratedNothing = iota
	FeaturesetuplinkV1710TxSupportUlGapfr2R17EnumeratedSupported
)
