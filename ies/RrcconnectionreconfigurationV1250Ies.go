package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1250-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1250Ies struct {
	WlanOffloadinfoR12   *interface{}
	ScgConfigurationR12  *ScgConfigurationR12
	SlSynctxcontrolR12   *SlSynctxcontrolR12
	SlDiscconfigR12      *SlDiscconfigR12
	SlCommconfigR12      *SlCommconfigR12
	Noncriticalextension *RrcconnectionreconfigurationV1310Ies
}
