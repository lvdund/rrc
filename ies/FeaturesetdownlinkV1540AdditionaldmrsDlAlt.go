package ies

import "rrc/utils"

// FeatureSetDownlink-v1540-additionalDMRS-DL-Alt ::= ENUMERATED
type FeaturesetdownlinkV1540AdditionaldmrsDlAlt struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1540AdditionaldmrsDlAltEnumeratedNothing = iota
	FeaturesetdownlinkV1540AdditionaldmrsDlAltEnumeratedSupported
)
