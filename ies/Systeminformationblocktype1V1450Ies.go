package ies

import "rrc/utils"

// SystemInformationBlockType1-v1450-IEs ::= SEQUENCE
type Systeminformationblocktype1V1450Ies struct {
	TddConfigV1450       *TddConfigV1450
	Noncriticalextension *Systeminformationblocktype1V1530Ies
}
