package ies

import "rrc/utils"

// SystemInformationBlockType1-v890-IEs ::= SEQUENCE
type Systeminformationblocktype1V890Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Systeminformationblocktype1V920Ies
}
