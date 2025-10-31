package ies

import "rrc/utils"

// AB-Config-NB-r13 ::= SEQUENCE
type AbConfigNbR13 struct {
	AbCategoryR13                AbConfigNbR13AbCategoryR13
	AbBarringbitmapR13           utils.BITSTRING `lb:10,ub:10`
	AbBarringforexceptiondataR13 *bool
	AbBarringforspecialacR13     utils.BITSTRING `lb:5,ub:5`
}
