package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-5GC-NB-r16-IEs ::= SEQUENCE
type Rrcconnectionreestablishmentrequest5gcNbR16Ies struct {
	UeIdentityR16           ReestabueIdentityCp5gcNbR16
	ReestablishmentcauseR16 ReestablishmentcauseNbR13
	CqiNpdcchR16            CqiNpdcchShortNbR14
	Spare                   utils.BITSTRING
}
