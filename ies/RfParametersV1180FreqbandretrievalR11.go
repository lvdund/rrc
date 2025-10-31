package ies

import "rrc/utils"

// RF-Parameters-v1180-freqBandRetrieval-r11 ::= ENUMERATED
type RfParametersV1180FreqbandretrievalR11 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1180FreqbandretrievalR11EnumeratedNothing = iota
	RfParametersV1180FreqbandretrievalR11EnumeratedSupported
)
