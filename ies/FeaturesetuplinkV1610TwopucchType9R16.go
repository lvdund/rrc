package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type9-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType9R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType9R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType9R16EnumeratedSupported
)
