package ies

import "rrc/utils"

// CondReconfigToAddMod-r16 ::= SEQUENCE
// Extensible
type CondreconfigtoaddmodR16 struct {
	CondreconfigidR16       CondreconfigidR16
	CondexecutioncondR16    *[]Measid `lb:1,ub:2`
	CondrrcreconfigR16      *utils.OCTETSTRING
	CondexecutioncondscgR17 *utils.OCTETSTRING `ext`
}
