package ies

import "rrc/utils"

// CG-COT-Sharing-r16-cot-Sharing-r16 ::= SEQUENCE
type CgCotSharingR16CotSharingR16 struct {
	DurationR16              utils.INTEGER `lb:0,ub:39`
	OffsetR16                utils.INTEGER `lb:0,ub:39`
	ChannelaccesspriorityR16 utils.INTEGER `lb:0,ub:4`
}
