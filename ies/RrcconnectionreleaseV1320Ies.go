package ies

import "rrc/utils"

// RRCConnectionRelease-v1320-IEs ::= SEQUENCE
type RrcconnectionreleaseV1320Ies struct {
	ResumeidentityR13    *ResumeidentityR13
	Noncriticalextension *RrcconnectionreleaseV1530Ies
}
