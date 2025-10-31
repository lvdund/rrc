package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type4-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType4R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType4R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType4R16EnumeratedSupported
)
