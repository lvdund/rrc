package ies

import "rrc/utils"

// PhysCellIdCDMA2000 ::= utils.INTEGER (0..maxPNOffset)
type Physcellidcdma2000 struct {
	Value utils.INTEGER `lb:0,ub:maxPNOffset`
}
