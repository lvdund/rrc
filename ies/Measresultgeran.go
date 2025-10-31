package ies

// MeasResultGERAN ::= SEQUENCE
// Extensible
type Measresultgeran struct {
	Carrierfreq Carrierfreqgeran
	Physcellid  Physcellidgeran
	CgiInfo     *MeasresultgeranCgiInfo
	Measresult  MeasresultgeranMeasresult
}
