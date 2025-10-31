package ies

import "rrc/utils"

// BandEntryIndex ::= utils.INTEGER (0.. maxNrofServingCells)
type Bandentryindex struct {
	Value utils.INTEGER `lb:0,ub:maxNrofServingCells`
}
