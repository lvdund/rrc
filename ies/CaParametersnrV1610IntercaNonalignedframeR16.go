package ies

import "rrc/utils"

// CA-ParametersNR-v1610-interCA-NonAlignedFrame-r16 ::= ENUMERATED
type CaParametersnrV1610IntercaNonalignedframeR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610IntercaNonalignedframeR16EnumeratedNothing = iota
	CaParametersnrV1610IntercaNonalignedframeR16EnumeratedSupported
)
