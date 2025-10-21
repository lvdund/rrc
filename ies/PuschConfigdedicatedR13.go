package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13 ::= SEQUENCE
type PuschConfigdedicatedR13 struct {
	BetaoffsetAckIndexR13   utils.INTEGER
	Betaoffset2AckIndexR13  *utils.INTEGER
	BetaoffsetRiIndexR13    utils.INTEGER
	BetaoffsetCqiIndexR13   utils.INTEGER
	BetaoffsetmcR13         *interface{}
	GrouphoppingdisabledR13 *utils.ENUMERATED
	DmrsWithoccActivatedR13 *utils.ENUMERATED
	PuschDmrsR11            *interface{}
	Ucionpusch              *interface{}
	PuschHoppingconfigR13   *utils.ENUMERATED
}
