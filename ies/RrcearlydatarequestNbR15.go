package ies

// RRCEarlyDataRequest-NB-r15-IEs ::= SEQUENCE
type RrcearlydatarequestNbR15 struct {
	STmsiR15              STmsi
	EstablishmentcauseR15 RrcearlydatarequestNbR15IesEstablishmentcauseR15
	CqiNpdcchR15          *CqiNpdcchNbR14
	DedicatedinfonasR15   Dedicatedinfonas
	Noncriticalextension  *RrcearlydatarequestNbV1590
}
