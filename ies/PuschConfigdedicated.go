package ies

import "rrc/utils"

// PUSCH-ConfigDedicated ::= SEQUENCE
type PuschConfigdedicated struct {
	BetaoffsetAckIndex utils.INTEGER `lb:0,ub:15`
	BetaoffsetRiIndex  utils.INTEGER `lb:0,ub:15`
	BetaoffsetCqiIndex utils.INTEGER `lb:0,ub:15`
}
