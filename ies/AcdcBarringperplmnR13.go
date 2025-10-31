package ies

import "rrc/utils"

// ACDC-BarringPerPLMN-r13 ::= SEQUENCE
type AcdcBarringperplmnR13 struct {
	PlmnIdentityindexR13          utils.INTEGER `lb:0,ub:maxPLMNR11`
	AcdcOnlyforhplmnR13           utils.BOOLEAN
	BarringperacdcCategorylistR13 BarringperacdcCategorylistR13
}
