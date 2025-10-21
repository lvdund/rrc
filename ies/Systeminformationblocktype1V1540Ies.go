package ies

import "rrc/utils"

// SystemInformationBlockType1-v1540-IEs ::= SEQUENCE
type Systeminformationblocktype1V1540Ies struct {
	SiPosoffsetR15       *utils.ENUMERATED
	Noncriticalextension *Systeminformationblocktype1V1610Ies
}
