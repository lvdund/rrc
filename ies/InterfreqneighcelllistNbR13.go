package ies

import "rrc/utils"

// InterFreqNeighCellList-NB-r13 ::= SEQUENCE OF PhysCellId
// SIZE (1..maxCellInter)
type InterfreqneighcelllistNbR13 struct {
	Value utils.Sequence[Physcellid]
}
