package ies

import "rrc/utils"

// PUCCH-format4 ::= SEQUENCE
type PucchFormat4 struct {
	Nrofsymbols         utils.INTEGER `lb:0,ub:14`
	OccLength           PucchFormat4OccLength
	OccIndex            PucchFormat4OccIndex
	Startingsymbolindex utils.INTEGER `lb:0,ub:10`
}
