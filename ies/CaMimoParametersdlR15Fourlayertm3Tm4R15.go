package ies

import "rrc/utils"

// CA-MIMO-ParametersDL-r15-fourLayerTM3-TM4-r15 ::= ENUMERATED
type CaMimoParametersdlR15Fourlayertm3Tm4R15 struct {
	Value utils.ENUMERATED
}

const (
	CaMimoParametersdlR15Fourlayertm3Tm4R15EnumeratedNothing = iota
	CaMimoParametersdlR15Fourlayertm3Tm4R15EnumeratedSupported
)
