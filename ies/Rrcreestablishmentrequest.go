package ies

import "rrc/utils"

// RRCReestablishmentRequest-IEs ::= SEQUENCE
type Rrcreestablishmentrequest struct {
	UeIdentity           ReestabueIdentity
	Reestablishmentcause Reestablishmentcause
	Spare                utils.BITSTRING `lb:1,ub:1`
}
