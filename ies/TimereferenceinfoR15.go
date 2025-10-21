package ies

import "rrc/utils"

// TimeReferenceInfo-r15 ::= SEQUENCE
type TimereferenceinfoR15 struct {
	TimeR15         ReferencetimeR15
	UncertaintyR15  *utils.INTEGER
	TimeinfotypeR15 *utils.ENUMERATED
	ReferencesfnR15 *utils.INTEGER
}
