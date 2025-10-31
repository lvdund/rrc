package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type2-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType2R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType2R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType2R16EnumeratedSupported
)
