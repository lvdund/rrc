package ies

import "rrc/utils"

// FeatureSetUplink-v1610-twoPUCCH-Type11-r16 ::= ENUMERATED
type FeaturesetuplinkV1610TwopucchType11R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610TwopucchType11R16EnumeratedNothing = iota
	FeaturesetuplinkV1610TwopucchType11R16EnumeratedSupported
)
