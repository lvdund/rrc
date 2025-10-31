package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-sfn-SchemeB-PDSCH-only-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700SfnSchemebPdschOnlyR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700SfnSchemebPdschOnlyR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700SfnSchemebPdschOnlyR17EnumeratedSupported
)
