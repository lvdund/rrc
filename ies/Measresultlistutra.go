package ies

// MeasResultListUTRA ::= SEQUENCE OF MeasResultUTRA
// SIZE (1..maxCellReport)
type Measresultlistutra struct {
	Value []Measresultutra `lb:1,ub:maxCellReport`
}
