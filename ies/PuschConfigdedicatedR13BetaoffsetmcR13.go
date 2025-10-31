package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13-betaOffsetMC-r13 ::= SEQUENCE
type PuschConfigdedicatedR13BetaoffsetmcR13 struct {
	BetaoffsetAckIndexMcR13  utils.INTEGER  `lb:0,ub:15`
	Betaoffset2AckIndexMcR13 *utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexMcR13   utils.INTEGER  `lb:0,ub:15`
	BetaoffsetCqiIndexMcR13  utils.INTEGER  `lb:0,ub:15`
}
