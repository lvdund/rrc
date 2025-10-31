package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-nPUCCH-Param-r13-setup ::= SEQUENCE
type PucchConfigdedicatedR13NpucchParamR13Setup struct {
	NpucchIdentityR13 utils.INTEGER `lb:0,ub:503`
	N1pucchAnR13      utils.INTEGER `lb:0,ub:2047`
}
