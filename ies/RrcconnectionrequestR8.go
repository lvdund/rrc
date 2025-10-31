package ies

import "rrc/utils"

// RRCConnectionRequest-r8-IEs ::= SEQUENCE
type RrcconnectionrequestR8 struct {
	UeIdentity         InitialueIdentity
	Establishmentcause Establishmentcause
	Spare              utils.BITSTRING `lb:1,ub:1`
}
