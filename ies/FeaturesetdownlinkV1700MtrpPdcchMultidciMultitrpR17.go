package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-mTRP-PDCCH-multiDCI-multiTRP-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700MtrpPdcchMultidciMultitrpR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700MtrpPdcchMultidciMultitrpR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700MtrpPdcchMultidciMultitrpR17EnumeratedSupported
)
