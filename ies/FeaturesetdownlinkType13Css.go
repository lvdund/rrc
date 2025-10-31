package ies

import "rrc/utils"

// FeatureSetDownlink-type1-3-CSS ::= ENUMERATED
type FeaturesetdownlinkType13Css struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkType13CssEnumeratedNothing = iota
	FeaturesetdownlinkType13CssEnumeratedSupported
)
