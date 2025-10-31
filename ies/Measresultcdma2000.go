package ies

// MeasResultCDMA2000 ::= SEQUENCE
// Extensible
type Measresultcdma2000 struct {
	Physcellid Physcellidcdma2000
	CgiInfo    *Cellglobalidcdma2000
	Measresult *Measresultcdma2000Measresult
}
