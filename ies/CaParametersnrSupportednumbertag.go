package ies

import "rrc/utils"

// CA-ParametersNR-supportedNumberTAG ::= ENUMERATED
type CaParametersnrSupportednumbertag struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrSupportednumbertagEnumeratedNothing = iota
	CaParametersnrSupportednumbertagEnumeratedN2
	CaParametersnrSupportednumbertagEnumeratedN3
	CaParametersnrSupportednumbertagEnumeratedN4
)
