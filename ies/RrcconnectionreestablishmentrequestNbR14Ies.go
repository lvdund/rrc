package ies

import "rrc/utils"

// RRCConnectionReestablishmentRequest-NB-r14-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrequestNbR14Ies struct {
	UeIdentityR14                ReestabueIdentityCpNbR14
	ReestablishmentcauseR14      ReestablishmentcauseNbR13
	CqiNpdcchR14                 CqiNpdcchShortNbR14
	EarlycontentionresolutionR14 bool
	Spare                        utils.BITSTRING
}
