package ies

import "rrc/utils"

// MAC-ParametersCommon-recommendedBitRateMultiplier-r16 ::= ENUMERATED
type MacParameterscommonRecommendedbitratemultiplierR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonRecommendedbitratemultiplierR16EnumeratedNothing = iota
	MacParameterscommonRecommendedbitratemultiplierR16EnumeratedSupported
)
