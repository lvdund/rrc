package ies

import "rrc/utils"

// ReselectionThreshold-NB-v1360 ::= utils.INTEGER (32..63)
type ReselectionthresholdNbV1360 struct {
	Value utils.INTEGER `lb:0,ub:63`
}
