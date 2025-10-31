package ies

import "rrc/utils"

// RACH-ConfigCommonTwoStepRA-r16-ra-PrioritizationForAccessIdentityTwoStep-r16 ::= SEQUENCE
type RachConfigcommontwostepraR16RaPrioritizationforaccessidentitytwostepR16 struct {
	RaPrioritizationR16      RaPrioritization
	RaPrioritizationforaiR16 utils.BITSTRING `lb:2,ub:2`
}
