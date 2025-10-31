package ies

// MeasResultListCDMA2000 ::= SEQUENCE OF MeasResultCDMA2000
// SIZE (1..maxCellReport)
type Measresultlistcdma2000 struct {
	Value []Measresultcdma2000 `lb:1,ub:maxCellReport`
}
