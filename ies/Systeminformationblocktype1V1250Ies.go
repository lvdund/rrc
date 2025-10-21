package ies

import "rrc/utils"

// SystemInformationBlockType1-v1250-IEs ::= SEQUENCE
type Systeminformationblocktype1V1250Ies struct {
	CellaccessrelatedinfoV1250   *interface{}
	CellselectioninfoV1250       *CellselectioninfoV1250
	FreqbandindicatorpriorityR12 *utils.ENUMERATED
	Noncriticalextension         *Systeminformationblocktype1V1310Ies
}
