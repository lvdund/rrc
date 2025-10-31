package ies

// RRCConnectionRelease-v1650-IEs ::= SEQUENCE
type RrcconnectionreleaseV1650 struct {
	MpspriorityindicationR16 *bool
	Noncriticalextension     *RrcconnectionreleaseV1650IesNoncriticalextension
}
