package ies

// RRCConnectionRelease-NB-v1550-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1550 struct {
	RedirectedcarrierinfoV1550 *RedirectedcarrierinfoNbV1550
	Noncriticalextension       *RrcconnectionreleaseNbV15b0
}
