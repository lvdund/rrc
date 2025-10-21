package ies

import "rrc/utils"

// SystemInformationBlockType1-v1350-IEs ::= SEQUENCE
type Systeminformationblocktype1V1350Ies struct {
	Cellselectioninfoce1R13 *Cellselectioninfoce1R13
	Noncriticalextension    *Systeminformationblocktype1V1360Ies
}
