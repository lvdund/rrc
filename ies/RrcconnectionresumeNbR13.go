package ies

import "rrc/utils"

// RRCConnectionResume-NB-r13-IEs ::= SEQUENCE
type RrcconnectionresumeNbR13 struct {
	RadioresourceconfigdedicatedR13 *RadioresourceconfigdedicatedNbR13
	NexthopchainingcountR13         Nexthopchainingcount
	DrbContinuerohcR13              *bool
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionresumeNbV1610
}
