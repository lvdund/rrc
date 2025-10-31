package ies

import "rrc/utils"

// CA-ParametersNRDC-v1700-simultaneousRxTx-IAB-MultipleParents-r17 ::= ENUMERATED
type CaParametersnrdcV1700SimultaneousrxtxIabMultipleparentsR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1700SimultaneousrxtxIabMultipleparentsR17EnumeratedNothing = iota
	CaParametersnrdcV1700SimultaneousrxtxIabMultipleparentsR17EnumeratedSupported
)
