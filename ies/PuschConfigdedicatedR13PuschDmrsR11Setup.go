package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13-pusch-DMRS-r11-setup ::= SEQUENCE
type PuschConfigdedicatedR13PuschDmrsR11Setup struct {
	NpuschIdentityR13   utils.INTEGER `lb:0,ub:509`
	NdmrsCshIdentityR13 utils.INTEGER `lb:0,ub:509`
}
