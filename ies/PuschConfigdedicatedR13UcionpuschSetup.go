package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13-uciOnPUSCH-setup ::= SEQUENCE
type PuschConfigdedicatedR13UcionpuschSetup struct {
	BetaoffsetAckIndexSubframeset2R13  utils.INTEGER  `lb:0,ub:15`
	Betaoffset2AckIndexSubframeset2R13 *utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndexSubframeset2R13   utils.INTEGER  `lb:0,ub:15`
	BetaoffsetCqiIndexSubframeset2R13  utils.INTEGER  `lb:0,ub:15`
	BetaoffsetmcR12                    *PuschConfigdedicatedR13UcionpuschSetupBetaoffsetmcR12
}
