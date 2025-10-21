package ies

import "rrc/utils"

// IntraFreqBlackCellList ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellBlack)
type Intrafreqblackcelllist struct {
	Value []Physcellidrange
}
