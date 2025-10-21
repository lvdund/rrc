package ies

import "rrc/utils"

// RRCConnectionReestablishment-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreestablishmentNbR13Ies struct {
	RadioresourceconfigdedicatedR13 RadioresourceconfigdedicatedNbR13
	NexthopchainingcountR13         Nexthopchainingcount
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionreestablishmentNbV1430Ies
}
