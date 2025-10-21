package ies

import "rrc/utils"

// RRCConnectionResume-v1530-IEs ::= SEQUENCE
type RrcconnectionresumeV1530Ies struct {
	FullconfigR15        *utils.ENUMERATED
	Noncriticalextension *RrcconnectionresumeV1610Ies
}
