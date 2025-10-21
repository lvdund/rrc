package ies

import "rrc/utils"

// SystemInformationBlockType1-v10x0-IEs ::= SEQUENCE
type Systeminformationblocktype1V10x0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Systeminformationblocktype1V12j0Ies
}
