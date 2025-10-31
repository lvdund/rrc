package ies

import "rrc/utils"

// CA-ParametersEUTRA-simultaneousRx-Tx ::= ENUMERATED
type CaParameterseutraSimultaneousrxTx struct {
	Value utils.ENUMERATED
}

const (
	CaParameterseutraSimultaneousrxTxEnumeratedNothing = iota
	CaParameterseutraSimultaneousrxTxEnumeratedSupported
)
