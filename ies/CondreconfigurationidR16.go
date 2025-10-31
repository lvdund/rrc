package ies

import "rrc/utils"

// CondReconfigurationId-r16 ::= utils.INTEGER (1.. maxCondConfig-r16)
type CondreconfigurationidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxCondConfigR16`
}
