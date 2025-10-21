package ies

import "rrc/utils"

// MeasResultListWLAN-r13 ::= SEQUENCE OF MeasResultWLAN-r13
// SIZE (1..maxCellReport)
type MeasresultlistwlanR13 struct {
	Value utils.Sequence[MeasresultwlanR13]
}
