package ies

import "rrc/utils"

// RRCConnectionRequest-5GC-r15-IEs ::= SEQUENCE
type Rrcconnectionrequest5gcR15 struct {
	UeIdentityR15         InitialueIdentity5gcR15
	EstablishmentcauseR15 Establishmentcause5gcR15
	Spare                 utils.BITSTRING `lb:1,ub:1`
}
