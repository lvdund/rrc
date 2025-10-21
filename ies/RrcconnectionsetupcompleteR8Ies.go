package ies

import "rrc/utils"

// RRCConnectionSetupComplete-r8-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteR8Ies struct {
	SelectedplmnIdentity utils.INTEGER
	Registeredmme        *Registeredmme
	Dedicatedinfonas     Dedicatedinfonas
	Noncriticalextension *RrcconnectionsetupcompleteV8a0Ies
}
