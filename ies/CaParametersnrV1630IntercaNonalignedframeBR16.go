package ies

import "rrc/utils"

// CA-ParametersNR-v1630-interCA-NonAlignedFrame-B-r16 ::= ENUMERATED
type CaParametersnrV1630IntercaNonalignedframeBR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1630IntercaNonalignedframeBR16EnumeratedNothing = iota
	CaParametersnrV1630IntercaNonalignedframeBR16EnumeratedSupported
)
