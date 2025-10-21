package ies

import "rrc/utils"

// RRCConnectionRequest-NB-r13-IEs ::= SEQUENCE
type RrcconnectionrequestNbR13Ies struct {
	UeIdentityR13                InitialueIdentity
	EstablishmentcauseR13        EstablishmentcauseNbR13
	MultitonesupportR13          *utils.ENUMERATED
	MulticarriersupportR13       *utils.ENUMERATED
	EarlycontentionresolutionR14 bool
	CqiNpdcchR14                 CqiNpdcchNbR14
	Spare                        utils.BITSTRING
}
