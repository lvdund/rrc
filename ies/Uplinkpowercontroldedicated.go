package ies

import "rrc/utils"

// UplinkPowerControlDedicated ::= SEQUENCE
type Uplinkpowercontroldedicated struct {
	P0UePusch           utils.INTEGER
	DeltamcsEnabled     utils.ENUMERATED
	Accumulationenabled bool
	P0UePucch           utils.INTEGER
	PsrsOffset          utils.INTEGER
	Filtercoefficient   Filtercoefficient
}
