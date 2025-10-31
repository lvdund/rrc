package ies

import "rrc/utils"

// P0-PUSCH-AlphaSet ::= SEQUENCE
type P0PuschAlphaset struct {
	P0PuschAlphasetid P0PuschAlphasetid
	P0                *utils.INTEGER `lb:0,ub:15`
	Alpha             *Alpha
}
