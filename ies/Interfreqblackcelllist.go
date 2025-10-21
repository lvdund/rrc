package ies

import "rrc/utils"

// InterFreqBlackCellList ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellBlack)
type Interfreqblackcelllist struct {
	Value utils.Sequence[Physcellidrange]
}
