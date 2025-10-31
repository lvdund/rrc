package ies

import "rrc/utils"

// MMTEL-Parameters-r14-recommendedBitRate-r14 ::= ENUMERATED
type MmtelParametersR14RecommendedbitrateR14 struct {
	Value utils.ENUMERATED
}

const (
	MmtelParametersR14RecommendedbitrateR14EnumeratedNothing = iota
	MmtelParametersR14RecommendedbitrateR14EnumeratedSupported
)
