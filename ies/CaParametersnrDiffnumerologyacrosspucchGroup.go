package ies

import "rrc/utils"

// CA-ParametersNR-diffNumerologyAcrossPUCCH-Group ::= ENUMERATED
type CaParametersnrDiffnumerologyacrosspucchGroup struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrDiffnumerologyacrosspucchGroupEnumeratedNothing = iota
	CaParametersnrDiffnumerologyacrosspucchGroupEnumeratedSupported
)
