package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-sfn-SchemeA-PDSCH-only-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700SfnSchemeaPdschOnlyR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700SfnSchemeaPdschOnlyR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700SfnSchemeaPdschOnlyR17EnumeratedSupported
)
