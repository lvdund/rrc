package ies

// MeasQuantityResultsEUTRA ::= SEQUENCE
type Measquantityresultseutra struct {
	Rsrp *RsrpRangeeutra
	Rsrq *RsrqRangeeutra
	Sinr *SinrRangeeutra
}
