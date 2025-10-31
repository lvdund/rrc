package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type5-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType5R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType5R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType5R16EnumeratedSupported
)
