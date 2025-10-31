package ies

import "rrc/utils"

// CondReconfigurationAddMod-r16 ::= SEQUENCE
// Extensible
type CondreconfigurationaddmodR16 struct {
	CondreconfigurationidR16      CondreconfigurationidR16
	TriggerconditionR16           *[]Measid `lb:1,ub:2`
	CondreconfigurationtoapplyR16 *utils.OCTETSTRING
}
