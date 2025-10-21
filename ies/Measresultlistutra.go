package ies

import "rrc/utils"

// MeasResultListUTRA ::= SEQUENCE OF MeasResultUTRA
// SIZE (1..maxCellReport)
type Measresultlistutra struct {
	Value utils.Sequence[Measresultutra]
}
