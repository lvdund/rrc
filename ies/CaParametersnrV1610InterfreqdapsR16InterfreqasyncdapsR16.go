package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interFreqDAPS-r16-interFreqAsyncDAPS-r16 ::= ENUMERATED
type CaParametersnrV1610InterfreqdapsR16InterfreqasyncdapsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610InterfreqdapsR16InterfreqasyncdapsR16EnumeratedNothing = iota
	CaParametersnrV1610InterfreqdapsR16InterfreqasyncdapsR16EnumeratedSupported
)
