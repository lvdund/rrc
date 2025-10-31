package ies

import "rrc/utils"

// RegisteredMME ::= SEQUENCE
type Registeredmme struct {
	PlmnIdentity *PlmnIdentity
	Mmegi        utils.BITSTRING `lb:16,ub:16`
	Mmec         Mmec
}
