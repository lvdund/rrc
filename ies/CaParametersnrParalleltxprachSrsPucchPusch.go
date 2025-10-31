package ies

import "rrc/utils"

// CA-ParametersNR-parallelTxPRACH-SRS-PUCCH-PUSCH ::= ENUMERATED
type CaParametersnrParalleltxprachSrsPucchPusch struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrParalleltxprachSrsPucchPuschEnumeratedNothing = iota
	CaParametersnrParalleltxprachSrsPucchPuschEnumeratedSupported
)
