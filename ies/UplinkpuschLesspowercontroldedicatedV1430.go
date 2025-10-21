package ies

import "rrc/utils"

// UplinkPUSCH-LessPowerControlDedicated-v1430 ::= SEQUENCE
type UplinkpuschLesspowercontroldedicatedV1430 struct {
	P0UePeriodicsrsR14     *utils.INTEGER
	P0UeAperiodicsrsR14    *utils.INTEGER
	AccumulationenabledR14 bool
}
