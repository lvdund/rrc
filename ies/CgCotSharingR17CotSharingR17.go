package ies

import "rrc/utils"

// CG-COT-Sharing-r17-cot-Sharing-r17 ::= SEQUENCE
type CgCotSharingR17CotSharingR17 struct {
	DurationR17 utils.INTEGER `lb:0,ub:319`
	OffsetR17   utils.INTEGER `lb:0,ub:319`
}
