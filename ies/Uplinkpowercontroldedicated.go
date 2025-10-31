package ies

import "rrc/utils"

// UplinkPowerControlDedicated ::= SEQUENCE
type Uplinkpowercontroldedicated struct {
	P0UePusch           utils.INTEGER `lb:0,ub:7`
	DeltamcsEnabled     UplinkpowercontroldedicatedDeltamcsEnabled
	Accumulationenabled utils.BOOLEAN
	P0UePucch           utils.INTEGER `lb:0,ub:7`
	PsrsOffset          utils.INTEGER `lb:0,ub:15`
	Filtercoefficient   Filtercoefficient
}
