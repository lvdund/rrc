package ies

import "rrc/utils"

// RRCConnectionRequest-NB-r13-IEs ::= SEQUENCE
type RrcconnectionrequestNbR13 struct {
	UeIdentityR13                InitialueIdentity
	EstablishmentcauseR13        EstablishmentcauseNbR13
	MultitonesupportR13          *bool
	MulticarriersupportR13       *bool
	EarlycontentionresolutionR14 utils.BOOLEAN
	CqiNpdcchR14                 CqiNpdcchNbR14
	Spare                        utils.BITSTRING `lb:17,ub:17`
}
