package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-r8-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrequestR8 struct {
	UeIdentity           ReestabueIdentity
	Reestablishmentcause Reestablishmentcause
	Spare                utils.BITSTRING `lb:2,ub:2`
}
