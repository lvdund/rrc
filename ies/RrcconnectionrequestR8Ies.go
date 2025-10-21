package ies

import "rrc/utils"

// RRCConnectionRequest-r8-IEs ::= SEQUENCE
type RrcconnectionrequestR8Ies struct {
	UeIdentity         InitialueIdentity
	Establishmentcause Establishmentcause
	Spare              utils.BITSTRING
}
