package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-nkaPUCCH-Param-r13-setup ::= SEQUENCE
type PucchConfigdedicatedR13NkapucchParamR13Setup struct {
	NkapucchAnR13 utils.INTEGER `lb:0,ub:2047`
}
