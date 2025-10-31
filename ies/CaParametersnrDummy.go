package ies

import "rrc/utils"

// CA-ParametersNR-dummy ::= ENUMERATED
type CaParametersnrDummy struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrDummyEnumeratedNothing = iota
	CaParametersnrDummyEnumeratedSupported
)
