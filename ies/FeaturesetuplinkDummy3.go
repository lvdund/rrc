package ies

import "rrc/utils"

// FeatureSetUplink-dummy3 ::= ENUMERATED
type FeaturesetuplinkDummy3 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkDummy3EnumeratedNothing = iota
	FeaturesetuplinkDummy3EnumeratedSupported
)
