package ies

import "rrc/utils"

// FeatureSetDownlink-v1540-oneFL-DMRS-ThreeAdditionalDMRS-DL ::= ENUMERATED
type FeaturesetdownlinkV1540OneflDmrsThreeadditionaldmrsDl struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1540OneflDmrsThreeadditionaldmrsDlEnumeratedNothing = iota
	FeaturesetdownlinkV1540OneflDmrsThreeadditionaldmrsDlEnumeratedSupported
)
