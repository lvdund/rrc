package ies

// RRCConnectionSetupComplete-NB-v1470-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbV1470 struct {
	MeasresultservcellR14 *MeasresultservcellNbR14
	Noncriticalextension  *RrcconnectionsetupcompleteNbV1610
}
