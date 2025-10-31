package ies

import "rrc/utils"

// CA-ParametersNR-v1560-diffNumerologyWithinPUCCH-GroupLargerSCS ::= ENUMERATED
type CaParametersnrV1560DiffnumerologywithinpucchGrouplargerscs struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1560DiffnumerologywithinpucchGrouplargerscsEnumeratedNothing = iota
	CaParametersnrV1560DiffnumerologywithinpucchGrouplargerscsEnumeratedSupported
)
