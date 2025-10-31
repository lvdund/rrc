package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrequestNbR13 struct {
	UeIdentityR13                ReestabueIdentity
	ReestablishmentcauseR13      ReestablishmentcauseNbR13
	CqiNpdcchR14                 CqiNpdcchNbR14
	EarlycontentionresolutionR14 utils.BOOLEAN
	Spare                        utils.BITSTRING `lb:20,ub:20`
}
