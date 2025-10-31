package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type8-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType8R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType8R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType8R16EnumeratedSupported
)
