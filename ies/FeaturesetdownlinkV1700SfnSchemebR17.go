package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-sfn-SchemeB-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700SfnSchemebR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700SfnSchemebR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700SfnSchemebR17EnumeratedSupported
)
