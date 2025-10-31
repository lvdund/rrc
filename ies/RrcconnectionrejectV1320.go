package ies

// RRCConnectionReject-v1320-IEs ::= SEQUENCE
type RrcconnectionrejectV1320 struct {
	RrcSuspendindicationR13 *bool
	Noncriticalextension    *RrcconnectionrejectV1320IesNoncriticalextension
}
