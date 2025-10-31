package ies

// RRCConnectionSetupComplete-v1610-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1610 struct {
	RlosRequestR16           *bool
	CpCiot5gsOptimisationR16 *bool
	UpCiot5gsOptimisationR16 *bool
	PurConfigidR16           *PurConfigidR16
	LteMR16                  *bool
	IabNodeindicationR16     *bool
	Noncriticalextension     *RrcconnectionsetupcompleteV1690
}
