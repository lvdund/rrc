package ies

import "rrc/utils"

// P0-PUCCH ::= SEQUENCE
type P0Pucch struct {
	P0PucchId    P0PucchId
	P0PucchValue utils.INTEGER `lb:0,ub:15`
}
