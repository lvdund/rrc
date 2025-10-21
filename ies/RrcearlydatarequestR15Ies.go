package ies

import "rrc/utils"

// RRCEarlyDataRequest-r15-IEs ::= SEQUENCE
type RrcearlydatarequestR15Ies struct {
	STmsiR15              STmsi
	EstablishmentcauseR15 utils.ENUMERATED
	DedicatedinfonasR15   Dedicatedinfonas
	Noncriticalextension  *RrcearlydatarequestV1590Ies
}
