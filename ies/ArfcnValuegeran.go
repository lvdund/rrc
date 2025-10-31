package ies

import "rrc/utils"

// ARFCN-ValueGERAN ::= utils.INTEGER (0..1023)
type ArfcnValuegeran struct {
	Value utils.INTEGER `lb:0,ub:1023`
}
