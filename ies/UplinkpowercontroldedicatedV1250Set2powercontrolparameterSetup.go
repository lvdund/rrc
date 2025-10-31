package ies

import "rrc/utils"

// UplinkPowerControlDedicated-v1250-set2PowerControlParameter-setup ::= SEQUENCE
type UplinkpowercontroldedicatedV1250Set2powercontrolparameterSetup struct {
	TpcSubframesetR12             utils.BITSTRING `lb:10,ub:10`
	P0NominalpuschSubframeset2R12 utils.INTEGER   `lb:0,ub:24`
	AlphaSubframeset2R12          AlphaR12
	P0UePuschSubframeset2R12      utils.INTEGER `lb:0,ub:7`
}
