package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbR13Ies struct {
	SelectedplmnIdentityR13         utils.INTEGER
	STmsiR13                        *STmsi
	RegisteredmmeR13                *Registeredmme
	DedicatedinfonasR13             Dedicatedinfonas
	AttachwithoutpdnConnectivityR13 *utils.ENUMERATED
	UpCiotEpsOptimisationR13        *utils.ENUMERATED
	Latenoncriticalextension        *utils.OCTETSTRING
	Noncriticalextension            *RrcconnectionsetupcompleteNbV1430Ies
}
