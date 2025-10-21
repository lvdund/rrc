package ies

import "rrc/utils"

// ReferenceTime-r15 ::= SEQUENCE
type ReferencetimeR15 struct {
	RefdaysR15                utils.INTEGER
	RefsecondsR15             utils.INTEGER
	RefmillisecondsR15        utils.INTEGER
	RefquartermicrosecondsR15 utils.INTEGER
}
