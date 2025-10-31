package ies

import "rrc/utils"

// FeatureSetDownlink-v1540-twoFL-DMRS-TwoAdditionalDMRS-DL ::= ENUMERATED
type FeaturesetdownlinkV1540TwoflDmrsTwoadditionaldmrsDl struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1540TwoflDmrsTwoadditionaldmrsDlEnumeratedNothing = iota
	FeaturesetdownlinkV1540TwoflDmrsTwoadditionaldmrsDlEnumeratedSupported
)
