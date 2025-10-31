package ies

import "rrc/utils"

// FeatureSetDownlink-v1610-singleDCI-SDM-scheme-r16 ::= ENUMERATED
type FeaturesetdownlinkV1610SingledciSdmSchemeR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1610SingledciSdmSchemeR16EnumeratedNothing = iota
	FeaturesetdownlinkV1610SingledciSdmSchemeR16EnumeratedSupported
)
