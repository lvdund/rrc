package ies

import "rrc/utils"

// ReselectionThreshold ::= utils.INTEGER (0..31)
type Reselectionthreshold struct {
	Value utils.INTEGER `lb:0,ub:31`
}
