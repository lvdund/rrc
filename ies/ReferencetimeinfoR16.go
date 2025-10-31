package ies

import "rrc/utils"

// ReferenceTimeInfo-r16 ::= SEQUENCE
type ReferencetimeinfoR16 struct {
	TimeR16         ReferencetimeR16
	UncertaintyR16  *utils.INTEGER `lb:0,ub:32767`
	TimeinfotypeR16 *ReferencetimeinfoR16TimeinfotypeR16
	ReferencesfnR16 *utils.INTEGER `lb:0,ub:1023`
}
