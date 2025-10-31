package ies

import "rrc/utils"

// CA-ParametersNR-simultaneousRxTxSUL ::= ENUMERATED
type CaParametersnrSimultaneousrxtxsul struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrSimultaneousrxtxsulEnumeratedNothing = iota
	CaParametersnrSimultaneousrxtxsulEnumeratedSupported
)
