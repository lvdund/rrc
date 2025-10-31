package ies

import "rrc/utils"

// VisitedCellInfo-r12 ::= SEQUENCE
// Extensible
type VisitedcellinfoR12 struct {
	VisitedcellidR12 *VisitedcellinfoR12VisitedcellidR12
	TimespentR12     utils.INTEGER `lb:0,ub:4095`
}
