package ies

import "rrc/utils"

// FeatureSetDownlink-v1540-pdsch-SeparationWithGap ::= ENUMERATED
type FeaturesetdownlinkV1540PdschSeparationwithgap struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1540PdschSeparationwithgapEnumeratedNothing = iota
	FeaturesetdownlinkV1540PdschSeparationwithgapEnumeratedSupported
)
