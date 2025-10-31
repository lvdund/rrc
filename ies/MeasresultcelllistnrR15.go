package ies

// MeasResultCellListNR-r15 ::= SEQUENCE OF MeasResultCellNR-r15
// SIZE (1..maxCellReport)
type MeasresultcelllistnrR15 struct {
	Value []MeasresultcellnrR15 `lb:1,ub:maxCellReport`
}
