package ies

import "rrc/utils"

// RRCEarlyDataRequest-NB-r15-IEs ::= SEQUENCE
type RrcearlydatarequestNbR15Ies struct {
	STmsiR15              STmsi
	EstablishmentcauseR15 utils.ENUMERATED
	CqiNpdcchR15          *CqiNpdcchNbR14
	DedicatedinfonasR15   Dedicatedinfonas
	Noncriticalextension  *RrcearlydatarequestNbV1590Ies
}
