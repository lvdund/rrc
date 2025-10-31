package ies

// RRCConnectionReestablishment-r8-IEs ::= SEQUENCE
type RrcconnectionreestablishmentR8 struct {
	Radioresourceconfigdedicated Radioresourceconfigdedicated
	Nexthopchainingcount         Nexthopchainingcount
	Noncriticalextension         *RrcconnectionreestablishmentV8a0
}
