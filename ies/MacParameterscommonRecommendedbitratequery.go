package ies

import "rrc/utils"

// MAC-ParametersCommon-recommendedBitRateQuery ::= ENUMERATED
type MacParameterscommonRecommendedbitratequery struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonRecommendedbitratequeryEnumeratedNothing = iota
	MacParameterscommonRecommendedbitratequeryEnumeratedSupported
)
