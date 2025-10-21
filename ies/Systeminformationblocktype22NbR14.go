package ies

import "rrc/utils"

// SystemInformationBlockType22-NB-r14 ::= SEQUENCE
// Extensible
type Systeminformationblocktype22NbR14 struct {
	DlConfiglistR14                *DlConfigcommonlistNbR14
	UlConfiglistR14                *UlConfigcommonlistNbR14
	PagingweightanchorR14          *PagingweightNbR14
	NprachProbabilityanchorlistR14 *NprachProbabilityanchorlistNbR14
	Latenoncriticalextension       *utils.OCTETSTRING
}
