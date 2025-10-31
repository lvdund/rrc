package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1250-uciOnPUSCH-setup ::= SEQUENCE
type PuschConfigdedicatedV1250UcionpuschSetup struct {
	BetaoffsetAckIndexSubframeset2R12 utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexSubframeset2R12  utils.INTEGER `lb:0,ub:15`
	BetaoffsetCqiIndexSubframeset2R12 utils.INTEGER `lb:0,ub:15`
	BetaoffsetmcR12                   *PuschConfigdedicatedV1250UcionpuschSetupBetaoffsetmcR12
}
