package ies

import "rrc/utils"

// ReferenceTime-r15 ::= SEQUENCE
type ReferencetimeR15 struct {
	RefdaysR15                utils.INTEGER `lb:0,ub:72999`
	RefsecondsR15             utils.INTEGER `lb:0,ub:86399`
	RefmillisecondsR15        utils.INTEGER `lb:0,ub:999`
	RefquartermicrosecondsR15 utils.INTEGER `lb:0,ub:3999`
}
