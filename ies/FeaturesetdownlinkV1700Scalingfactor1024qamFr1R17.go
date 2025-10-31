package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-scalingFactor-1024QAM-FR1-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700Scalingfactor1024qamFr1R17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700Scalingfactor1024qamFr1R17EnumeratedNothing = iota
	FeaturesetdownlinkV1700Scalingfactor1024qamFr1R17EnumeratedF0p4
	FeaturesetdownlinkV1700Scalingfactor1024qamFr1R17EnumeratedF0p75
	FeaturesetdownlinkV1700Scalingfactor1024qamFr1R17EnumeratedF0p8
)
