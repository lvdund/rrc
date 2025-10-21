package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-r8-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrequestR8Ies struct {
	UeIdentity           ReestabueIdentity
	Reestablishmentcause Reestablishmentcause
	Spare                utils.BITSTRING
}
