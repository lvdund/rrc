package ies

import "rrc/utils"

// FeatureSetUplink-v1610-dummy1 ::= ENUMERATED
type FeaturesetuplinkV1610Dummy1 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610Dummy1EnumeratedNothing = iota
	FeaturesetuplinkV1610Dummy1EnumeratedSupported
)
