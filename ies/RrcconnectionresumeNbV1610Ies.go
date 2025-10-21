package ies

import "rrc/utils"

// RRCConnectionResume-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionresumeNbV1610Ies struct {
	FullconfigR16        *utils.ENUMERATED
	Noncriticalextension *interface{}
}
