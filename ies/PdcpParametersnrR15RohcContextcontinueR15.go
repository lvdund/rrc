package ies

import "rrc/utils"

// PDCP-ParametersNR-r15-rohc-ContextContinue-r15 ::= ENUMERATED
type PdcpParametersnrR15RohcContextcontinueR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersnrR15RohcContextcontinueR15EnumeratedNothing = iota
	PdcpParametersnrR15RohcContextcontinueR15EnumeratedSupported
)
