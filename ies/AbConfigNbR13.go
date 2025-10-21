package ies

import "rrc/utils"

// AB-Config-NB-r13 ::= SEQUENCE
type AbConfigNbR13 struct {
	AbCategoryR13                utils.ENUMERATED
	AbBarringbitmapR13           utils.BITSTRING
	AbBarringforexceptiondataR13 *utils.ENUMERATED
	AbBarringforspecialacR13     utils.BITSTRING
}
