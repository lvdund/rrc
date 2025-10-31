package ies

import "rrc/utils"

// FeatureSetUplink-v1630-dummy ::= ENUMERATED
type FeaturesetuplinkV1630Dummy struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1630DummyEnumeratedNothing = iota
	FeaturesetuplinkV1630DummyEnumeratedSupported
)
