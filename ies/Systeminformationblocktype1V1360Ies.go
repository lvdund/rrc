package ies

import "rrc/utils"

// SystemInformationBlockType1-v1360-IEs ::= SEQUENCE
type Systeminformationblocktype1V1360Ies struct {
	Cellselectioninfoce1V1360 *Cellselectioninfoce1V1360
	Noncriticalextension      *Systeminformationblocktype1V1430Ies
}
