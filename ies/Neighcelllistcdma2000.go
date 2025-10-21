package ies

import "rrc/utils"

// NeighCellListCDMA2000 ::= SEQUENCE OF NeighCellCDMA2000
// SIZE (1..16)
type Neighcelllistcdma2000 struct {
	Value utils.Sequence[Neighcellcdma2000]
}
