package ies

import "rrc/utils"

// CA-ParametersNR-v1640-diffNumerologyWithinPUCCH-GroupLargerSCS-CarrierTypes-r16 ::= ENUMERATED
type CaParametersnrV1640DiffnumerologywithinpucchGrouplargerscsCarriertypesR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1640DiffnumerologywithinpucchGrouplargerscsCarriertypesR16EnumeratedNothing = iota
	CaParametersnrV1640DiffnumerologywithinpucchGrouplargerscsCarriertypesR16EnumeratedSupported
)
