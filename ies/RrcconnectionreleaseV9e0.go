package ies

// RRCConnectionRelease-v9e0-IEs ::= SEQUENCE
type RrcconnectionreleaseV9e0 struct {
	RedirectedcarrierinfoV9e0       *RedirectedcarrierinfoV9e0
	IdlemodemobilitycontrolinfoV9e0 *IdlemodemobilitycontrolinfoV9e0
	Noncriticalextension            *RrcconnectionreleaseV9e0IesNoncriticalextension
}
