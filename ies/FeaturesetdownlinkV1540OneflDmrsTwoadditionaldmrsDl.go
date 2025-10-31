package ies

import "rrc/utils"

// FeatureSetDownlink-v1540-oneFL-DMRS-TwoAdditionalDMRS-DL ::= ENUMERATED
type FeaturesetdownlinkV1540OneflDmrsTwoadditionaldmrsDl struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1540OneflDmrsTwoadditionaldmrsDlEnumeratedNothing = iota
	FeaturesetdownlinkV1540OneflDmrsTwoadditionaldmrsDlEnumeratedSupported
)
