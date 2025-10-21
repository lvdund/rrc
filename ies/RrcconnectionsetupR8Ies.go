package ies

import "rrc/utils"

// RRCConnectionSetup-r8-IEs ::= SEQUENCE
type RrcconnectionsetupR8Ies struct {
	Radioresourceconfigdedicated Radioresourceconfigdedicated
	Noncriticalextension         *RrcconnectionsetupV8a0Ies
}
