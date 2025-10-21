package ies

import "rrc/utils"

// PUSCH-ConfigDedicated ::= SEQUENCE
type PuschConfigdedicated struct {
	BetaoffsetAckIndex utils.INTEGER
	BetaoffsetRiIndex  utils.INTEGER
	BetaoffsetCqiIndex utils.INTEGER
}
