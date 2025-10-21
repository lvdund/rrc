package ies

import "rrc/utils"

// MeasResultServCellListSCG-r12 ::= SEQUENCE OF MeasResultServCellSCG-r12
// SIZE (1..maxServCell-r10)
type MeasresultservcelllistscgR12 struct {
	Value []MeasresultservcellscgR12
}
