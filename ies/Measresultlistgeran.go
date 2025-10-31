package ies

// MeasResultListGERAN ::= SEQUENCE OF MeasResultGERAN
// SIZE (1..maxCellReport)
type Measresultlistgeran struct {
	Value []Measresultgeran `lb:1,ub:maxCellReport`
}
