package ies

import "rrc/utils"

// CA-ParametersEUTRA-additionalRx-Tx-PerformanceReq ::= ENUMERATED
type CaParameterseutraAdditionalrxTxPerformancereq struct {
	Value utils.ENUMERATED
}

const (
	CaParameterseutraAdditionalrxTxPerformancereqEnumeratedNothing = iota
	CaParameterseutraAdditionalrxTxPerformancereqEnumeratedSupported
)
