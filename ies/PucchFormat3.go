package ies

import "rrc/utils"

// PUCCH-format3 ::= SEQUENCE
type PucchFormat3 struct {
	Nrofprbs            utils.INTEGER `lb:0,ub:16`
	Nrofsymbols         utils.INTEGER `lb:0,ub:14`
	Startingsymbolindex utils.INTEGER `lb:0,ub:10`
}
