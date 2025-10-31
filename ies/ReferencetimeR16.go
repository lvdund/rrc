package ies

import "rrc/utils"

// ReferenceTime-r16 ::= SEQUENCE
type ReferencetimeR16 struct {
	RefdaysR16           utils.INTEGER `lb:0,ub:72999`
	RefsecondsR16        utils.INTEGER `lb:0,ub:86399`
	RefmillisecondsR16   utils.INTEGER `lb:0,ub:999`
	ReftennanosecondsR16 utils.INTEGER `lb:0,ub:99999`
}
