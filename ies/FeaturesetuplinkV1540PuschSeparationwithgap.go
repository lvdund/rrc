package ies

import "rrc/utils"

// FeatureSetUplink-v1540-pusch-SeparationWithGap ::= ENUMERATED
type FeaturesetuplinkV1540PuschSeparationwithgap struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1540PuschSeparationwithgapEnumeratedNothing = iota
	FeaturesetuplinkV1540PuschSeparationwithgapEnumeratedSupported
)
