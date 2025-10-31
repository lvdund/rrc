package ies

import "rrc/utils"

// P0AlphaSet-r17-closedLoopIndex-r17 ::= ENUMERATED
type P0alphasetR17ClosedloopindexR17 struct {
	Value utils.ENUMERATED
}

const (
	P0alphasetR17ClosedloopindexR17EnumeratedNothing = iota
	P0alphasetR17ClosedloopindexR17EnumeratedI0
	P0alphasetR17ClosedloopindexR17EnumeratedI1
)
