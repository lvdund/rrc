package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type3-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType3R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType3R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType3R16EnumeratedSupported
)
