package ies

import "rrc/utils"

// CA-ParametersNR-parallelTxSRS-PUCCH-PUSCH ::= ENUMERATED
type CaParametersnrParalleltxsrsPucchPusch struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrParalleltxsrsPucchPuschEnumeratedNothing = iota
	CaParametersnrParalleltxsrsPucchPuschEnumeratedSupported
)
