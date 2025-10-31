package ies

// CondReconfigurationToAddModList-r16 ::= SEQUENCE OF CondReconfigurationAddMod-r16
// SIZE (1.. maxCondConfig-r16)
type CondreconfigurationtoaddmodlistR16 struct {
	Value []CondreconfigurationaddmodR16 `lb:1,ub:maxCondConfigR16`
}
