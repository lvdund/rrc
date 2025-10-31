package ies

import "rrc/utils"

// CA-ParametersNR-v1610-simul-SRS-MIMO-Trans-BC-r16 ::= ENUMERATED
type CaParametersnrV1610SimulSrsMimoTransBcR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610SimulSrsMimoTransBcR16EnumeratedNothing = iota
	CaParametersnrV1610SimulSrsMimoTransBcR16EnumeratedN2
)
