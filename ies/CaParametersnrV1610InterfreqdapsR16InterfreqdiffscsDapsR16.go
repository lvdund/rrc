package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interFreqDAPS-r16-interFreqDiffSCS-DAPS-r16 ::= ENUMERATED
type CaParametersnrV1610InterfreqdapsR16InterfreqdiffscsDapsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610InterfreqdapsR16InterfreqdiffscsDapsR16EnumeratedNothing = iota
	CaParametersnrV1610InterfreqdapsR16InterfreqdiffscsDapsR16EnumeratedSupported
)
