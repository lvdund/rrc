package ies

// RRCConnectionReconfiguration-r8-IEs ::= SEQUENCE
type RrcconnectionreconfigurationR8 struct {
	Measconfig                   *Measconfig
	Mobilitycontrolinfo          *Mobilitycontrolinfo
	Dedicatedinfonaslist         *[]Dedicatedinfonas `lb:1,ub:maxDRB`
	Radioresourceconfigdedicated *Radioresourceconfigdedicated
	Securityconfigho             *Securityconfigho
	Noncriticalextension         *RrcconnectionreconfigurationV890
}
