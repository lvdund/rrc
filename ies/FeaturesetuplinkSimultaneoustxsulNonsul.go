package ies

import "rrc/utils"

// FeatureSetUplink-simultaneousTxSUL-NonSUL ::= ENUMERATED
type FeaturesetuplinkSimultaneoustxsulNonsul struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkSimultaneoustxsulNonsulEnumeratedNothing = iota
	FeaturesetuplinkSimultaneoustxsulNonsulEnumeratedSupported
)
