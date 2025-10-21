package ies

import "rrc/utils"

// SystemInformationBlockType1-v920-IEs ::= SEQUENCE
type Systeminformationblocktype1V920Ies struct {
	ImsEmergencysupportR9 *utils.ENUMERATED
	CellselectioninfoV920 *CellselectioninfoV920
	Noncriticalextension  *Systeminformationblocktype1V1130Ies
}
