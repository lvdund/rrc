package ies

import "rrc/utils"

// IRAT-ParametersNR-r15-en-DC-r15 ::= ENUMERATED
type IratParametersnrR15EnDcR15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrR15EnDcR15EnumeratedNothing = iota
	IratParametersnrR15EnDcR15EnumeratedSupported
)
