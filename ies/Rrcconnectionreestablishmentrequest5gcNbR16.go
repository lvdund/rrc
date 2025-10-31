package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-5GC-NB-r16-IEs ::= SEQUENCE
type Rrcconnectionreestablishmentrequest5gcNbR16 struct {
	UeIdentityR16           ReestabueIdentityCp5gcNbR16
	ReestablishmentcauseR16 ReestablishmentcauseNbR13
	CqiNpdcchR16            CqiNpdcchShortNbR14
	Spare                   utils.BITSTRING `lb:1,ub:1`
}
