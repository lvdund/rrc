package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-sfn-SchemeA-PDCCH-only-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700SfnSchemeaPdcchOnlyR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700SfnSchemeaPdcchOnlyR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700SfnSchemeaPdcchOnlyR17EnumeratedSupported
)
