package ies

import "rrc/utils"

// CA-ParametersNR-v1550-dummy ::= ENUMERATED
type CaParametersnrV1550Dummy struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1550DummyEnumeratedNothing = iota
	CaParametersnrV1550DummyEnumeratedSupported
)
