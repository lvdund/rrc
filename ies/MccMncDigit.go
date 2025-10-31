package ies

import "rrc/utils"

// MCC-MNC-Digit ::= utils.INTEGER (0..9)
type MccMncDigit struct {
	Value utils.INTEGER `lb:0,ub:9`
}
