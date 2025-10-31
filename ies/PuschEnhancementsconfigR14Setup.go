package ies

import "rrc/utils"

// PUSCH-EnhancementsConfig-r14-setup ::= SEQUENCE
type PuschEnhancementsconfigR14Setup struct {
	PuschHoppingoffsetpuschEnhR14 *utils.INTEGER `lb:0,ub:100`
	IntervalUlhoppingpuschEnhR14  *PuschEnhancementsconfigR14SetupIntervalUlhoppingpuschEnhR14
}
