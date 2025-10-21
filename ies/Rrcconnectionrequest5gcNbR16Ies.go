package ies

import "rrc/utils"

// RRCConnectionRequest-5GC-NB-r16-IEs ::= SEQUENCE
type Rrcconnectionrequest5gcNbR16Ies struct {
	UeIdentityR16         InitialueIdentity5gcNbR16
	EstablishmentcauseR16 utils.ENUMERATED
	CqiNpdcchR16          CqiNpdcchNbR14
	Spare                 utils.BITSTRING
}
