package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1470-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbV1470Ies struct {
	MeasresultservcellR14 *MeasresultservcellNbR14
	Noncriticalextension  *RrcconnectionsetupcompleteNbV1610Ies
}
