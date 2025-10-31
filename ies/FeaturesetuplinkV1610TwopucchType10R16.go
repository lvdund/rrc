package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type10-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType10R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType10R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType10R16EnumeratedSupported
)
