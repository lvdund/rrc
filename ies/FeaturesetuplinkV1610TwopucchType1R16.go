package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type1-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType1R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType1R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType1R16EnumeratedSupported
)
