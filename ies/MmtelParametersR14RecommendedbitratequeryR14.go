package ies

import "rrc/utils"

// MMTEL-Parameters-r14-recommendedBitRateQuery-r14 ::= ENUMERATED
type MmtelParametersR14RecommendedbitratequeryR14 struct {
	Value utils.ENUMERATED
}

const (
	MmtelParametersR14RecommendedbitratequeryR14EnumeratedNothing = iota
	MmtelParametersR14RecommendedbitratequeryR14EnumeratedSupported
)
