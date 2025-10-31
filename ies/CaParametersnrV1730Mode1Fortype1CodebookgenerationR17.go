package ies

import "rrc/utils"

// CA-ParametersNR-v1730-mode1-ForType1-CodebookGeneration-r17 ::= ENUMERATED
type CaParametersnrV1730Mode1Fortype1CodebookgenerationR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730Mode1Fortype1CodebookgenerationR17EnumeratedNothing = iota
	CaParametersnrV1730Mode1Fortype1CodebookgenerationR17EnumeratedSupported
)
