package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type6-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType6R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType6R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType6R16EnumeratedSupported
)
