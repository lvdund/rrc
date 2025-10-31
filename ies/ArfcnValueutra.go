package ies

import "rrc/utils"

// ARFCN-ValueUTRA ::= utils.INTEGER (0..16383)
type ArfcnValueutra struct {
	Value utils.INTEGER `lb:0,ub:16383`
}
