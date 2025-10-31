package ies

import "rrc/utils"

// FeatureSetDownlink-dummy8 ::= ENUMERATED
type FeaturesetdownlinkDummy8 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkDummy8EnumeratedNothing = iota
	FeaturesetdownlinkDummy8EnumeratedSupported
)
