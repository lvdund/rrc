package ies

import "rrc/utils"

// RRCConnectionReconfiguration-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreconfigurationNbR13 struct {
	DedicatedinfonaslistR13         *[]Dedicatedinfonas `lb:1,ub:maxDRBNbR13`
	RadioresourceconfigdedicatedR13 *RadioresourceconfigdedicatedNbR13
	FullconfigR13                   *bool
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionreconfigurationNbR13IesNoncriticalextension
}
