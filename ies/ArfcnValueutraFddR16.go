package ies

import "rrc/utils"

// ARFCN-ValueUTRA-FDD-r16 ::= utils.INTEGER (0..16383)
type ArfcnValueutraFddR16 struct {
	Value utils.INTEGER `lb:0,ub:16383`
}
