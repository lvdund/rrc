package ies

import "rrc/utils"

// SL-MinT2Value-r15 ::= SEQUENCE
type SlMint2valueR15 struct {
	PrioritylistR15 SlPrioritylistR13
	Mint2valueR15   utils.INTEGER `lb:0,ub:20`
}
