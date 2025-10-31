package ies

// MeasResultListUTRA-FDD-r16 ::= SEQUENCE OF MeasResultUTRA-FDD-r16
// SIZE (1..maxCellReport)
type MeasresultlistutraFddR16 struct {
	Value []MeasresultutraFddR16 `lb:1,ub:maxCellReport`
}
