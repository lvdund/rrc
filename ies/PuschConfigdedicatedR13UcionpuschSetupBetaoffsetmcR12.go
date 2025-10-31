package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13-uciOnPUSCH-setup-betaOffsetMC-r12 ::= SEQUENCE
type PuschConfigdedicatedR13UcionpuschSetupBetaoffsetmcR12 struct {
	BetaoffsetAckIndexMcSubframeset2R13  utils.INTEGER  `lb:0,ub:15`
	Betaoffset2AckIndexMcSubframeset2R13 *utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexMcSubframeset2R13   utils.INTEGER  `lb:0,ub:15`
	BetaoffsetCqiIndexMcSubframeset2R13  utils.INTEGER  `lb:0,ub:15`
}
