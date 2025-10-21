package ies

import "rrc/utils"

// MeasResultListCDMA2000 ::= SEQUENCE OF MeasResultCDMA2000
// SIZE (1..maxCellReport)
type Measresultlistcdma2000 struct {
	Value utils.Sequence[Measresultcdma2000]
}
