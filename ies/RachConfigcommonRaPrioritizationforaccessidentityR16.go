package ies

import "rrc/utils"

// RACH-ConfigCommon-ra-PrioritizationForAccessIdentity-r16 ::= SEQUENCE
type RachConfigcommonRaPrioritizationforaccessidentityR16 struct {
	RaPrioritizationR16      RaPrioritization
	RaPrioritizationforaiR16 utils.BITSTRING `lb:2,ub:2`
}
