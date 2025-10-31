package ies

import "rrc/utils"

// CA-ParametersNR-simultaneousRxTxInterBandCA ::= ENUMERATED
type CaParametersnrSimultaneousrxtxinterbandca struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrSimultaneousrxtxinterbandcaEnumeratedNothing = iota
	CaParametersnrSimultaneousrxtxinterbandcaEnumeratedSupported
)
