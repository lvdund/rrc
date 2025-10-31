package ies

import "rrc/utils"

// MAC-ParametersCommon-recommendedBitRate ::= ENUMERATED
type MacParameterscommonRecommendedbitrate struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonRecommendedbitrateEnumeratedNothing = iota
	MacParameterscommonRecommendedbitrateEnumeratedSupported
)
