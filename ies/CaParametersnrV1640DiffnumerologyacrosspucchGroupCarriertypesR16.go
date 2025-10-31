package ies

import "rrc/utils"

// CA-ParametersNR-v1640-diffNumerologyAcrossPUCCH-Group-CarrierTypes-r16 ::= ENUMERATED
type CaParametersnrV1640DiffnumerologyacrosspucchGroupCarriertypesR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1640DiffnumerologyacrosspucchGroupCarriertypesR16EnumeratedNothing = iota
	CaParametersnrV1640DiffnumerologyacrosspucchGroupCarriertypesR16EnumeratedSupported
)
