package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1320-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1320Ies struct {
	CeModebR13                      *utils.ENUMERATED
	STmsiR13                        *STmsi
	AttachwithoutpdnConnectivityR13 *utils.ENUMERATED
	UpCiotEpsOptimisationR13        *utils.ENUMERATED
	CpCiotEpsOptimisationR13        *utils.ENUMERATED
	Noncriticalextension            *RrcconnectionsetupcompleteV1330Ies
}
