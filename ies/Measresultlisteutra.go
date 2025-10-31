package ies

// MeasResultListEUTRA ::= SEQUENCE OF MeasResultEUTRA
// SIZE (1..maxCellReport)
type Measresultlisteutra struct {
	Value []Measresulteutra `lb:1,ub:maxCellReport`
}
