package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbR13 struct {
	SelectedplmnIdentityR13         utils.INTEGER `lb:0,ub:maxPLMNR11`
	STmsiR13                        *STmsi
	RegisteredmmeR13                *Registeredmme
	DedicatedinfonasR13             Dedicatedinfonas
	AttachwithoutpdnConnectivityR13 *bool
	UpCiotEpsOptimisationR13        *bool
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionsetupcompleteNbV1430
}
