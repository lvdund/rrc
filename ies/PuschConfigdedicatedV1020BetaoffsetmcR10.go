package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1020-betaOffsetMC-r10 ::= SEQUENCE
type PuschConfigdedicatedV1020BetaoffsetmcR10 struct {
	BetaoffsetAckIndexMcR10 utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexMcR10  utils.INTEGER `lb:0,ub:15`
	BetaoffsetCqiIndexMcR10 utils.INTEGER `lb:0,ub:15`
}
