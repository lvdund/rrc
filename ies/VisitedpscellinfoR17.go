package ies

import "rrc/utils"

// VisitedPSCellInfo-r17 ::= SEQUENCE
// Extensible
type VisitedpscellinfoR17 struct {
	VisitedcellidR17 *VisitedpscellinfoR17VisitedcellidR17
	TimespentR17     utils.INTEGER `lb:0,ub:4095`
}
