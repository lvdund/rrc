package ies

import "rrc/utils"

// PhysCellIdListCDMA2000 ::= SEQUENCE OF PhysCellIdCDMA2000
// SIZE (1..16)
type Physcellidlistcdma2000 struct {
	Value utils.Sequence[Physcellidcdma2000]
}
