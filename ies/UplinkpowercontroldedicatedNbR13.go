package ies

import "rrc/utils"

// UplinkPowerControlDedicated-NB-r13 ::= SEQUENCE
type UplinkpowercontroldedicatedNbR13 struct {
	P0UeNpuschR13 utils.INTEGER `lb:0,ub:7`
}
