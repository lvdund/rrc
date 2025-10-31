package ies

import "rrc/utils"

// TimeReferenceInfo-r15 ::= SEQUENCE
type TimereferenceinfoR15 struct {
	TimeR15         ReferencetimeR15
	UncertaintyR15  *utils.INTEGER `lb:0,ub:12`
	TimeinfotypeR15 *TimereferenceinfoR15TimeinfotypeR15
	ReferencesfnR15 *utils.INTEGER `lb:0,ub:1023`
}
