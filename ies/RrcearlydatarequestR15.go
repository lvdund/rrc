package ies

// RRCEarlyDataRequest-r15-IEs ::= SEQUENCE
type RrcearlydatarequestR15 struct {
	STmsiR15              STmsi
	EstablishmentcauseR15 RrcearlydatarequestR15IesEstablishmentcauseR15
	DedicatedinfonasR15   Dedicatedinfonas
	Noncriticalextension  *RrcearlydatarequestV1590
}
