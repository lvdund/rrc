package ies

import "rrc/utils"

// RRCConnectionRequest-5GC-NB-r16-IEs ::= SEQUENCE
type Rrcconnectionrequest5gcNbR16 struct {
	UeIdentityR16         InitialueIdentity5gcNbR16
	EstablishmentcauseR16 Rrcconnectionrequest5gcNbR16IesEstablishmentcauseR16
	CqiNpdcchR16          CqiNpdcchNbR14
	Spare                 utils.BITSTRING `lb:11,ub:11`
}
