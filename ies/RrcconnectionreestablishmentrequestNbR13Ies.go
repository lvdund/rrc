package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrequestNbR13Ies struct {
	UeIdentityR13                ReestabueIdentity
	ReestablishmentcauseR13      ReestablishmentcauseNbR13
	CqiNpdcchR14                 CqiNpdcchNbR14
	EarlycontentionresolutionR14 bool
	Spare                        utils.BITSTRING
}
