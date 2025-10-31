package ies

// RRCConnectionReconfiguration-v1250-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1250 struct {
	WlanOffloadinfoR12   *RrcconnectionreconfigurationV1250IesWlanOffloadinfoR12
	ScgConfigurationR12  *ScgConfigurationR12
	SlSynctxcontrolR12   *SlSynctxcontrolR12
	SlDiscconfigR12      *SlDiscconfigR12
	SlCommconfigR12      *SlCommconfigR12
	Noncriticalextension *RrcconnectionreconfigurationV1310
}
