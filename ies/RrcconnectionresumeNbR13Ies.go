package ies

import "rrc/utils"

// RRCConnectionResume-NB-r13-IEs ::= SEQUENCE
type RrcconnectionresumeNbR13Ies struct {
	RadioresourceconfigdedicatedR13 *RadioresourceconfigdedicatedNbR13
	NexthopchainingcountR13         Nexthopchainingcount
	DrbContinuerohcR13              *utils.ENUMERATED
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionresumeNbV1610Ies
}
