package ies

import "rrc/utils"

// PUCCH-format2 ::= SEQUENCE
type PucchFormat2 struct {
	Nrofprbs            utils.INTEGER `lb:0,ub:16`
	Nrofsymbols         utils.INTEGER `lb:0,ub:2`
	Startingsymbolindex utils.INTEGER `lb:0,ub:13`
}
