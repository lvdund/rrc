package ies

import "rrc/utils"

// RRCConnectionSetupComplete-r8-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteR8 struct {
	SelectedplmnIdentity utils.INTEGER `lb:0,ub:maxPLMNR11`
	Registeredmme        *Registeredmme
	Dedicatedinfonas     Dedicatedinfonas
	Noncriticalextension *RrcconnectionsetupcompleteV8a0
}
