package ies

import "rrc/utils"

// BLER-Result-r12 ::= SEQUENCE
type BlerResultR12 struct {
	BlerR12           BlerRangeR12
	BlocksreceivedR12 interface{}
}
