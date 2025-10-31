package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interFreqDAPS-r16-interFreqMultiUL-TransmissionDAPS-r16 ::= ENUMERATED
type CaParametersnrV1610InterfreqdapsR16InterfreqmultiulTransmissiondapsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610InterfreqdapsR16InterfreqmultiulTransmissiondapsR16EnumeratedNothing = iota
	CaParametersnrV1610InterfreqdapsR16InterfreqmultiulTransmissiondapsR16EnumeratedSupported
)
