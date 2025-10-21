package ies

import "rrc/utils"

// NeighCellListCDMA2000-v920 ::= SEQUENCE OF NeighCellCDMA2000-v920
// SIZE (1..16)
type Neighcelllistcdma2000V920 struct {
	Value []Neighcellcdma2000V920
}
