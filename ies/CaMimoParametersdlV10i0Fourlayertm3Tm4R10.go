package ies

import "rrc/utils"

// CA-MIMO-ParametersDL-v10i0-fourLayerTM3-TM4-r10 ::= ENUMERATED
type CaMimoParametersdlV10i0Fourlayertm3Tm4R10 struct {
	Value utils.ENUMERATED
}

const (
	CaMimoParametersdlV10i0Fourlayertm3Tm4R10EnumeratedNothing = iota
	CaMimoParametersdlV10i0Fourlayertm3Tm4R10EnumeratedSupported
)
