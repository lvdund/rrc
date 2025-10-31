package ies

// CondReconfigurationToRemoveList-r16 ::= SEQUENCE OF CondReconfigurationId-r16
// SIZE (1..maxCondConfig-r16)
type CondreconfigurationtoremovelistR16 struct {
	Value []CondreconfigurationidR16 `lb:1,ub:maxCondConfigR16`
}
