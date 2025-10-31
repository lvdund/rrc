package ies

import "rrc/utils"

// CA-MIMO-ParametersDL-r13-fourLayerTM3-TM4-r13 ::= ENUMERATED
type CaMimoParametersdlR13Fourlayertm3Tm4R13 struct {
	Value utils.ENUMERATED
}

const (
	CaMimoParametersdlR13Fourlayertm3Tm4R13EnumeratedNothing = iota
	CaMimoParametersdlR13Fourlayertm3Tm4R13EnumeratedSupported
)
