package ies

// MeasResultListRSSI-SCG-r13 ::= SEQUENCE OF MeasResultRSSI-SCG-r13
// SIZE (1..maxServCell-r13)
type MeasresultlistrssiScgR13 struct {
	Value []MeasresultrssiScgR13 `lb:1,ub:maxServCellR13`
}
