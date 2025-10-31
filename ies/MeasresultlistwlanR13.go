package ies

// MeasResultListWLAN-r13 ::= SEQUENCE OF MeasResultWLAN-r13
// SIZE (1..maxCellReport)
type MeasresultlistwlanR13 struct {
	Value []MeasresultwlanR13 `lb:1,ub:maxCellReport`
}
