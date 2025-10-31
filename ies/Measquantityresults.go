package ies

// MeasQuantityResults ::= SEQUENCE
type Measquantityresults struct {
	Rsrp *RsrpRange
	Rsrq *RsrqRange
	Sinr *SinrRange
}
