package ies

// MeasResultCBR-r14 ::= SEQUENCE
type MeasresultcbrR14 struct {
	PoolidentityR14 SlV2xTxpoolreportidentityR14
	CbrPsschR14     SlCbrR14
	CbrPscchR14     *SlCbrR14
}
