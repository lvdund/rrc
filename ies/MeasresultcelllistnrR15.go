package ies

import "rrc/utils"

// MeasResultCellListNR-r15 ::= SEQUENCE OF MeasResultCellNR-r15
// SIZE (1..maxCellReport)
type MeasresultcelllistnrR15 struct {
	Value utils.Sequence[MeasresultcellnrR15]
}
