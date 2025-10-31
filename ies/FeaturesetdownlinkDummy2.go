package ies

import "rrc/utils"

// FeatureSetDownlink-dummy2 ::= ENUMERATED
type FeaturesetdownlinkDummy2 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkDummy2EnumeratedNothing = iota
	FeaturesetdownlinkDummy2EnumeratedSupported
)
