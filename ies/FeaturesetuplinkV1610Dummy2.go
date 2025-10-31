package ies

import "rrc/utils"

// FeatureSetUplink-v1610-dummy2 ::= ENUMERATED
type FeaturesetuplinkV1610Dummy2 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610Dummy2EnumeratedNothing = iota
	FeaturesetuplinkV1610Dummy2EnumeratedSupported
)
