package ies

import "rrc/utils"

// UplinkPowerControlDedicated-v1530 ::= SEQUENCE
type UplinkpowercontroldedicatedV1530 struct {
	AlphaUeR15   *AlphaR12
	P0UePuschR15 *utils.INTEGER `lb:0,ub:15`
}
