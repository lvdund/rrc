package ies

// RRCConnectionSetup-r8-IEs ::= SEQUENCE
type RrcconnectionsetupR8 struct {
	Radioresourceconfigdedicated Radioresourceconfigdedicated
	Noncriticalextension         *RrcconnectionsetupV8a0
}
