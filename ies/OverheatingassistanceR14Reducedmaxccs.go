package ies

import "rrc/utils"

// OverheatingAssistance-r14-reducedMaxCCs ::= SEQUENCE
type OverheatingassistanceR14Reducedmaxccs struct {
	Reducedccsdl utils.INTEGER `lb:0,ub:31`
	Reducedccsul utils.INTEGER `lb:0,ub:31`
}
