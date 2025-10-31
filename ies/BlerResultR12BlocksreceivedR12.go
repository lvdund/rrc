package ies

import "rrc/utils"

// BLER-Result-r12-blocksReceived-r12 ::= SEQUENCE
type BlerResultR12BlocksreceivedR12 struct {
	NR12 utils.BITSTRING `lb:3,ub:3`
	MR12 utils.BITSTRING `lb:8,ub:8`
}
