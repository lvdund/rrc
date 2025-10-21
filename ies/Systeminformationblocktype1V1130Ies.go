package ies

import "rrc/utils"

// SystemInformationBlockType1-v1130-IEs ::= SEQUENCE
type Systeminformationblocktype1V1130Ies struct {
	TddConfigV1130         *TddConfigV1130
	CellselectioninfoV1130 *CellselectioninfoV1130
	Noncriticalextension   *Systeminformationblocktype1V1250Ies
}
