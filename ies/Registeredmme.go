package ies

import "rrc/utils"

// RegisteredMME ::= SEQUENCE
type Registeredmme struct {
	PlmnIdentity *PlmnIdentity
	Mmegi        utils.BITSTRING
	Mmec         Mmec
}
