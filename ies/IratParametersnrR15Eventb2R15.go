package ies

import "rrc/utils"

// IRAT-ParametersNR-r15-eventB2-r15 ::= ENUMERATED
type IratParametersnrR15Eventb2R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrR15Eventb2R15EnumeratedNothing = iota
	IratParametersnrR15Eventb2R15EnumeratedSupported
)
