package ies

import "rrc/utils"

// PUCCH-format1 ::= SEQUENCE
type PucchFormat1 struct {
	Initialcyclicshift  utils.INTEGER `lb:0,ub:11`
	Nrofsymbols         utils.INTEGER `lb:0,ub:14`
	Startingsymbolindex utils.INTEGER `lb:0,ub:10`
	Timedomainocc       utils.INTEGER `lb:0,ub:6`
}
