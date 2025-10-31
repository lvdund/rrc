package ies

import "rrc/utils"

// FeatureSetUplink-v1540-zeroSlotOffsetAperiodicSRS ::= ENUMERATED
type FeaturesetuplinkV1540Zeroslotoffsetaperiodicsrs struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1540ZeroslotoffsetaperiodicsrsEnumeratedNothing = iota
	FeaturesetuplinkV1540ZeroslotoffsetaperiodicsrsEnumeratedSupported
)
