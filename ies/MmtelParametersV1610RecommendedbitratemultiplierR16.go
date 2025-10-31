package ies

import "rrc/utils"

// MMTEL-Parameters-v1610-recommendedBitRateMultiplier-r16 ::= ENUMERATED
type MmtelParametersV1610RecommendedbitratemultiplierR16 struct {
	Value utils.ENUMERATED
}

const (
	MmtelParametersV1610RecommendedbitratemultiplierR16EnumeratedNothing = iota
	MmtelParametersV1610RecommendedbitratemultiplierR16EnumeratedSupported
)
