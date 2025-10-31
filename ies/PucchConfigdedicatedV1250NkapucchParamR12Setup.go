package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1250-nkaPUCCH-Param-r12-setup ::= SEQUENCE
type PucchConfigdedicatedV1250NkapucchParamR12Setup struct {
	NkapucchAnR12 utils.INTEGER `lb:0,ub:2047`
}
