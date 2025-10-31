package ies

import "rrc/utils"

// IRAT-ParametersNR-v1660-extendedBand-n77-r16 ::= ENUMERATED
type IratParametersnrV1660ExtendedbandN77R16 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1660ExtendedbandN77R16EnumeratedNothing = iota
	IratParametersnrV1660ExtendedbandN77R16EnumeratedSupported
)
