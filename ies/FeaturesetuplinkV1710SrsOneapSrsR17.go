package ies

import "rrc/utils"

// FeatureSetUplink-v1710-srs-OneAP-SRS-r17 ::= ENUMERATED
type FeaturesetuplinkV1710SrsOneapSrsR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710SrsOneapSrsR17EnumeratedNothing = iota
	FeaturesetuplinkV1710SrsOneapSrsR17EnumeratedSupported
)
