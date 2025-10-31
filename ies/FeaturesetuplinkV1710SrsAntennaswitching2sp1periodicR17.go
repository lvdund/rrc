package ies

import "rrc/utils"

// FeatureSetUplink-v1710-srs-AntennaSwitching2SP-1Periodic-r17 ::= ENUMERATED
type FeaturesetuplinkV1710SrsAntennaswitching2sp1periodicR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710SrsAntennaswitching2sp1periodicR17EnumeratedNothing = iota
	FeaturesetuplinkV1710SrsAntennaswitching2sp1periodicR17EnumeratedSupported
)
