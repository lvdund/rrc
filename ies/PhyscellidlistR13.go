package ies

import "rrc/utils"

// PhysCellIdList-r13 ::= SEQUENCE OF PhysCellId
// SIZE (1.. maxSL-DiscCells-r13)
type PhyscellidlistR13 struct {
	Value utils.Sequence[Physcellid]
}
