package ies

// MeasResultCellListSFTD-EUTRA ::= SEQUENCE OF MeasResultSFTD-EUTRA
// SIZE (1..maxCellSFTD)
type MeasresultcelllistsftdEutra struct {
	Value []MeasresultsftdEutra `lb:1,ub:maxCellSFTD`
}
