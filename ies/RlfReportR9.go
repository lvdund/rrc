package ies

// RLF-Report-r9 ::= SEQUENCE
// Extensible
type RlfReportR9 struct {
	MeasresultlastservcellR9 *RlfReportR9MeasresultlastservcellR9
	MeasresultneighcellsR9   *RlfReportR9MeasresultneighcellsR9
}
