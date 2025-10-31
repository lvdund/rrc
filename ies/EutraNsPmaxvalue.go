package ies

import "rrc/utils"

// EUTRA-NS-PmaxValue ::= SEQUENCE
type EutraNsPmaxvalue struct {
	Additionalpmax             *utils.INTEGER `lb:0,ub:33`
	Additionalspectrumemission *utils.INTEGER `lb:0,ub:288`
}
