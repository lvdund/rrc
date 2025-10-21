package ies

import "rrc/utils"

// RLF-Report-r9 ::= SEQUENCE
// Extensible
type RlfReportR9 struct {
	MeasresultlastservcellR9 *interface{}
	MeasresultneighcellsR9   *interface{}
}
