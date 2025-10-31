package ies

import "rrc/utils"

// CA-ParametersNRDC-v1610-asyncNRDC-r16 ::= ENUMERATED
type CaParametersnrdcV1610AsyncnrdcR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1610AsyncnrdcR16EnumeratedNothing = iota
	CaParametersnrdcV1610AsyncnrdcR16EnumeratedSupported
)
