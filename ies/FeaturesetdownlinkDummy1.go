package ies

import "rrc/utils"

// FeatureSetDownlink-dummy1 ::= ENUMERATED
type FeaturesetdownlinkDummy1 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkDummy1EnumeratedNothing = iota
	FeaturesetdownlinkDummy1EnumeratedSupported
)
