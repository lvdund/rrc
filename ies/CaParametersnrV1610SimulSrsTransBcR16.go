package ies

import "rrc/utils"

// CA-ParametersNR-v1610-simul-SRS-Trans-BC-r16 ::= ENUMERATED
type CaParametersnrV1610SimulSrsTransBcR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610SimulSrsTransBcR16EnumeratedNothing = iota
	CaParametersnrV1610SimulSrsTransBcR16EnumeratedN2
)
