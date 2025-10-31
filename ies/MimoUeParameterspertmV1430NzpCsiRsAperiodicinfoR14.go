package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-nzp-CSI-RS-AperiodicInfo-r14 ::= SEQUENCE
type MimoUeParameterspertmV1430NzpCsiRsAperiodicinfoR14 struct {
	NmaxprocR14     utils.INTEGER `lb:0,ub:32`
	NmaxresourceR14 MimoUeParameterspertmV1430NzpCsiRsAperiodicinfoR14NmaxresourceR14
}
