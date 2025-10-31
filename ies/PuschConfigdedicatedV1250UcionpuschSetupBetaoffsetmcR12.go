package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1250-uciOnPUSCH-setup-betaOffsetMC-r12 ::= SEQUENCE
type PuschConfigdedicatedV1250UcionpuschSetupBetaoffsetmcR12 struct {
	BetaoffsetAckIndexMcSubframeset2R12 utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexMcSubframeset2R12  utils.INTEGER `lb:0,ub:15`
	BetaoffsetCqiIndexMcSubframeset2R12 utils.INTEGER `lb:0,ub:15`
}
