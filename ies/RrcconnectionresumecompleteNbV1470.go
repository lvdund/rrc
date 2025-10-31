package ies

// RRCConnectionResumeComplete-NB-v1470-IEs ::= SEQUENCE
type RrcconnectionresumecompleteNbV1470 struct {
	MeasresultservcellR14 *MeasresultservcellNbR14
	Noncriticalextension  *RrcconnectionresumecompleteNbV1610
}
