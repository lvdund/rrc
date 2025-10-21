package ies

import "rrc/utils"

// RRCConnectionReconfiguration-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreconfigurationNbR13Ies struct {
	DedicatedinfonaslistR13         *interface{}
	RadioresourceconfigdedicatedR13 *RadioresourceconfigdedicatedNbR13
	FullconfigR13                   *utils.ENUMERATED
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *interface{}
}
