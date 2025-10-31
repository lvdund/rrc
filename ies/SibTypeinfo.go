package ies

import "rrc/utils"

// SIB-TypeInfo ::= SEQUENCE
// Extensible
type SibTypeinfo struct {
	Type      SibTypeinfoType
	Valuetag  *utils.INTEGER `lb:0,ub:31`
	Areascope *utils.BOOLEAN
}
