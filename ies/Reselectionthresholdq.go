package ies

import "rrc/utils"

// ReselectionThresholdQ ::= utils.INTEGER (0..31)
type Reselectionthresholdq struct {
	Value utils.INTEGER `lb:0,ub:31`
}
