package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1130-pusch-DMRS-r11-setup ::= SEQUENCE
type PuschConfigdedicatedV1130PuschDmrsR11Setup struct {
	NpuschIdentityR11   utils.INTEGER `lb:0,ub:509`
	NdmrsCshIdentityR11 utils.INTEGER `lb:0,ub:509`
}
