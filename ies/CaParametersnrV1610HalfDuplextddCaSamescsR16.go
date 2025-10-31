package ies

import "rrc/utils"

// CA-ParametersNR-v1610-half-DuplexTDD-CA-SameSCS-r16 ::= ENUMERATED
type CaParametersnrV1610HalfDuplextddCaSamescsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610HalfDuplextddCaSamescsR16EnumeratedNothing = iota
	CaParametersnrV1610HalfDuplextddCaSamescsR16EnumeratedSupported
)
