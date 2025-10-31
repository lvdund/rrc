package ies

// MeasResultCellListSFTD-r15 ::= SEQUENCE OF MeasResultCellSFTD-r15
// SIZE (1..maxCellSFTD)
type MeasresultcelllistsftdR15 struct {
	Value []MeasresultcellsftdR15 `lb:1,ub:maxCellSFTD`
}
