package ies

import "rrc/utils"

// CA-ParametersNR-v1640-diffNumerologyWithinPUCCH-GroupSmallerSCS-CarrierTypes-r16 ::= ENUMERATED
type CaParametersnrV1640DiffnumerologywithinpucchGroupsmallerscsCarriertypesR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1640DiffnumerologywithinpucchGroupsmallerscsCarriertypesR16EnumeratedNothing = iota
	CaParametersnrV1640DiffnumerologywithinpucchGroupsmallerscsCarriertypesR16EnumeratedSupported
)
