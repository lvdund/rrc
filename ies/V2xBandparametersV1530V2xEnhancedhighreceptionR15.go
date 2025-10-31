package ies

import "rrc/utils"

// V2X-BandParameters-v1530-v2x-EnhancedHighReception-r15 ::= ENUMERATED
type V2xBandparametersV1530V2xEnhancedhighreceptionR15 struct {
	Value utils.ENUMERATED
}

const (
	V2xBandparametersV1530V2xEnhancedhighreceptionR15EnumeratedNothing = iota
	V2xBandparametersV1530V2xEnhancedhighreceptionR15EnumeratedSupported
)
