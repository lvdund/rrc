package ies

import "rrc/utils"

// IMSI-Digit ::= utils.INTEGER (0..9)
type ImsiDigit struct {
	Value utils.INTEGER `lb:0,ub:9`
}
