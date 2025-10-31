package ies

import "rrc/utils"

// RF-Parameters-v1250-freqBandPriorityAdjustment-r12 ::= ENUMERATED
type RfParametersV1250FreqbandpriorityadjustmentR12 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1250FreqbandpriorityadjustmentR12EnumeratedNothing = iota
	RfParametersV1250FreqbandpriorityadjustmentR12EnumeratedSupported
)
