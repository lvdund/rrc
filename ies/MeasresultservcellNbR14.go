package ies

import "rrc/utils"

// MeasResultServCell-NB-r14 ::= SEQUENCE
type MeasresultservcellNbR14 struct {
	NrsrpresultR14 NrsrpRangeNbR14
	NrsrqresultR14 NrsrqRangeNbR14
}
