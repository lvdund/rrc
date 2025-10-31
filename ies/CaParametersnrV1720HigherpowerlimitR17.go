package ies

import "rrc/utils"

// CA-ParametersNR-v1720-higherPowerLimit-r17 ::= ENUMERATED
type CaParametersnrV1720HigherpowerlimitR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720HigherpowerlimitR17EnumeratedNothing = iota
	CaParametersnrV1720HigherpowerlimitR17EnumeratedSupported
)
