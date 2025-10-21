package ies

import "rrc/utils"

// SystemInformationBlockType1-NB-v1350 ::= SEQUENCE
type Systeminformationblocktype1NbV1350 struct {
	CellselectioninfoV1350 *CellselectioninfoNbV1350
	Noncriticalextension   *Systeminformationblocktype1NbV1430
}
