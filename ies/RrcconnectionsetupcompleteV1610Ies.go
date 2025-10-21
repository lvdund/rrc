package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1610-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1610Ies struct {
	RlosRequestR16           *utils.ENUMERATED
	CpCiot5gsOptimisationR16 *utils.ENUMERATED
	UpCiot5gsOptimisationR16 *utils.ENUMERATED
	PurConfigidR16           *PurConfigidR16
	LteMR16                  *utils.ENUMERATED
	IabNodeindicationR16     *utils.ENUMERATED
	Noncriticalextension     *RrcconnectionsetupcompleteV1690Ies
}
