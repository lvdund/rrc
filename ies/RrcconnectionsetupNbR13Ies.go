package ies

import "rrc/utils"

// RRCConnectionSetup-NB-r13-IEs ::= SEQUENCE
type RrcconnectionsetupNbR13Ies struct {
	RadioresourceconfigdedicatedR13 RadioresourceconfigdedicatedNbR13
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionsetupNbV1610Ies
}
