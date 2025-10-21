package ies

import "rrc/utils"

// MeasGapConfigToRemoveList-r14 ::= SEQUENCE OF ServCellIndex-r13
// SIZE (1..maxServCell-r13)
type MeasgapconfigtoremovelistR14 struct {
	Value utils.Sequence[ServcellindexR13]
}
