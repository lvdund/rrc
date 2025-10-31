package ies

import "rrc/utils"

// ReselectionThresholdQ-r9 ::= utils.INTEGER (0..31)
type ReselectionthresholdqR9 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
