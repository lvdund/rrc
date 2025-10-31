package ies

import "rrc/utils"

// UE-RadioPagingInfo-r12 ::= SEQUENCE
// Extensible
type UeRadiopaginginfoR12 struct {
	UeCategoryV1250 *utils.INTEGER `lb:0,ub:0`
}
