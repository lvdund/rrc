package ies

import "rrc/utils"

// UplinkPUSCH-LessPowerControlDedicated-v1430 ::= SEQUENCE
type UplinkpuschLesspowercontroldedicatedV1430 struct {
	P0UePeriodicsrsR14     *utils.INTEGER `lb:0,ub:7`
	P0UeAperiodicsrsR14    *utils.INTEGER `lb:0,ub:7`
	AccumulationenabledR14 utils.BOOLEAN
}
