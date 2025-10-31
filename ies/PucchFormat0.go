package ies

import "rrc/utils"

// PUCCH-format0 ::= SEQUENCE
type PucchFormat0 struct {
	Initialcyclicshift  utils.INTEGER `lb:0,ub:11`
	Nrofsymbols         utils.INTEGER `lb:0,ub:2`
	Startingsymbolindex utils.INTEGER `lb:0,ub:13`
}
