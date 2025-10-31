package ies

// MeasResultCellListSFTD-NR ::= SEQUENCE OF MeasResultCellSFTD-NR
// SIZE (1..maxCellSFTD)
type MeasresultcelllistsftdNr struct {
	Value []MeasresultcellsftdNr `lb:1,ub:maxCellSFTD`
}
