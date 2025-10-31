package ies

// MeasResultList2GERAN-r10 ::= SEQUENCE OF MeasResultListGERAN
// SIZE (1..maxCellListGERAN)
type Measresultlist2geranR10 struct {
	Value []Measresultlistgeran `lb:1,ub:maxCellListGERAN`
}
