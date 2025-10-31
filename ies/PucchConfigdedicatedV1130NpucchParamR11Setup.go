package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1130-nPUCCH-Param-r11-setup ::= SEQUENCE
type PucchConfigdedicatedV1130NpucchParamR11Setup struct {
	NpucchIdentityR11 utils.INTEGER `lb:0,ub:503`
	N1pucchAnR11      utils.INTEGER `lb:0,ub:2047`
}
