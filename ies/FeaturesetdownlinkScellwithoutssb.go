package ies

import "rrc/utils"

// FeatureSetDownlink-scellWithoutSSB ::= ENUMERATED
type FeaturesetdownlinkScellwithoutssb struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkScellwithoutssbEnumeratedNothing = iota
	FeaturesetdownlinkScellwithoutssbEnumeratedSupported
)
