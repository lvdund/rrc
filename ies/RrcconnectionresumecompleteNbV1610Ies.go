package ies

import "rrc/utils"

// RRCConnectionResumeComplete-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionresumecompleteNbV1610Ies struct {
	RlfInfoavailableR16  *utils.ENUMERATED
	AnrInfoavailableR16  *utils.ENUMERATED
	Noncriticalextension *interface{}
}
