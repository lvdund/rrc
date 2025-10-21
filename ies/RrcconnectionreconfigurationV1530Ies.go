package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1530-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1530Ies struct {
	SecurityconfighoV1530      *SecurityconfighoV1530
	ScellgrouptoreleaselistR15 *ScellgrouptoreleaselistR15
	ScellgrouptoaddmodlistR15  *ScellgrouptoaddmodlistR15
	DedicatedinfonaslistR15    *interface{}
	PMaxueFr1R15               *PMax
	SmtcR15                    *MtcSsbNrR15
	Noncriticalextension       *RrcconnectionreconfigurationV1610Ies
}
