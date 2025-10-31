package ies

import "rrc/utils"

// OverheatingAssistance-r14-reducedUE-Category ::= SEQUENCE
type OverheatingassistanceR14ReducedueCategory struct {
	ReducedueCategorydl utils.INTEGER `lb:0,ub:19`
	ReducedueCategoryul utils.INTEGER `lb:0,ub:21`
}
