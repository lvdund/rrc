package ies

import "rrc/utils"

// RRCConnectionReconfiguration-r8-IEs ::= SEQUENCE
type RrcconnectionreconfigurationR8Ies struct {
	Measconfig                   *Measconfig
	Mobilitycontrolinfo          *Mobilitycontrolinfo
	Dedicatedinfonaslist         *interface{}
	Radioresourceconfigdedicated *Radioresourceconfigdedicated
	Securityconfigho             *Securityconfigho
	Noncriticalextension         *RrcconnectionreconfigurationV890Ies
}
