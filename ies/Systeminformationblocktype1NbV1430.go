package ies

import "rrc/utils"

// SystemInformationBlockType1-NB-v1430 ::= SEQUENCE
type Systeminformationblocktype1NbV1430 struct {
	CellselectioninfoV1430 *CellselectioninfoNbV1430
	Noncriticalextension   *Systeminformationblocktype1NbV1450
}
