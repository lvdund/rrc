package ies

import "rrc/utils"

// RRCSetupRequest-IEs ::= SEQUENCE
type Rrcsetuprequest struct {
	UeIdentity         InitialueIdentity
	Establishmentcause Establishmentcause
	Spare              utils.BITSTRING `lb:1,ub:1`
}
