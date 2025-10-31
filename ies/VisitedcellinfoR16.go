package ies

import "rrc/utils"

// VisitedCellInfo-r16 ::= SEQUENCE
// Extensible
type VisitedcellinfoR16 struct {
	VisitedcellidR16               *VisitedcellinfoR16VisitedcellidR16
	TimespentR16                   utils.INTEGER             `lb:0,ub:4095`
	VisitedpscellinfolistreportR17 *VisitedpscellinfolistR17 `ext`
}
