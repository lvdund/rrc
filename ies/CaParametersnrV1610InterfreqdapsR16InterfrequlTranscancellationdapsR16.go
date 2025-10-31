package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interFreqDAPS-r16-interFreqUL-TransCancellationDAPS-r16 ::= ENUMERATED
type CaParametersnrV1610InterfreqdapsR16InterfrequlTranscancellationdapsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610InterfreqdapsR16InterfrequlTranscancellationdapsR16EnumeratedNothing = iota
	CaParametersnrV1610InterfreqdapsR16InterfrequlTranscancellationdapsR16EnumeratedSupported
)
