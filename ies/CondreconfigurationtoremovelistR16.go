package ies

import "rrc/utils"

// CondReconfigurationToRemoveList-r16 ::= SEQUENCE OF CondReconfigurationId-r16
// SIZE (1..maxCondConfig-r16)
type CondreconfigurationtoremovelistR16 struct {
	Value utils.Sequence[CondreconfigurationidR16]
}
