package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13 ::= SEQUENCE
type PuschConfigdedicatedR13 struct {
	BetaoffsetAckIndexR13   utils.INTEGER  `lb:0,ub:15`
	Betaoffset2AckIndexR13  *utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexR13    utils.INTEGER  `lb:0,ub:15`
	BetaoffsetCqiIndexR13   utils.INTEGER  `lb:0,ub:15`
	BetaoffsetmcR13         *PuschConfigdedicatedR13BetaoffsetmcR13
	GrouphoppingdisabledR13 *bool
	DmrsWithoccActivatedR13 *bool
	PuschDmrsR11            *PuschConfigdedicatedR13PuschDmrsR11
	Ucionpusch              *PuschConfigdedicatedR13Ucionpusch
	PuschHoppingconfigR13   *PuschConfigdedicatedR13PuschHoppingconfigR13
}
