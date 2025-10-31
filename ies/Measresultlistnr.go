package ies

// MeasResultListNR ::= SEQUENCE OF MeasResultNR
// SIZE (1..maxCellReport)
type Measresultlistnr struct {
	Value []Measresultnr `lb:1,ub:maxCellReport`
}
