package ies

import "rrc/utils"

// CA-ParametersNR-diffNumerologyWithinPUCCH-GroupSmallerSCS ::= ENUMERATED
type CaParametersnrDiffnumerologywithinpucchGroupsmallerscs struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrDiffnumerologywithinpucchGroupsmallerscsEnumeratedNothing = iota
	CaParametersnrDiffnumerologywithinpucchGroupsmallerscsEnumeratedSupported
)
