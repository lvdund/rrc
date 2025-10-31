package ies

// RRCConnectionSetupComplete-v1320-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1320 struct {
	CeModebR13                      *RrcconnectionsetupcompleteV1320IesCeModebR13
	STmsiR13                        *STmsi
	AttachwithoutpdnConnectivityR13 *bool
	UpCiotEpsOptimisationR13        *bool
	CpCiotEpsOptimisationR13        *bool
	Noncriticalextension            *RrcconnectionsetupcompleteV1330
}
