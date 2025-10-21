package ies

import "rrc/utils"

// MeasResultListGERAN ::= SEQUENCE OF MeasResultGERAN
// SIZE (1..maxCellReport)
type Measresultlistgeran struct {
	Value utils.Sequence[Measresultgeran]
}
