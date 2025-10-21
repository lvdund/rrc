package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-NB-v1470-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteNbV1470Ies struct {
	MeasresultservcellR14 *MeasresultservcellNbR14
	Noncriticalextension  *RrcconnectionreestablishmentcompleteNbV1610Ies
}
