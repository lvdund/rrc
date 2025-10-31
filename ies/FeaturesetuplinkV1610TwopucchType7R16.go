package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type7-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType7R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType7R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType7R16EnumeratedSupported
)
