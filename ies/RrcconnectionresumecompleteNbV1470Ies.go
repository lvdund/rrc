package ies

import "rrc/utils"

// RRCConnectionResumeComplete-NB-v1470-IEs ::= SEQUENCE
type RrcconnectionresumecompleteNbV1470Ies struct {
	MeasresultservcellR14 *MeasresultservcellNbR14
	Noncriticalextension  *RrcconnectionresumecompleteNbV1610Ies
}
