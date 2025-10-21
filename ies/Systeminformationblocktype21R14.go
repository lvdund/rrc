package ies

import "rrc/utils"

// SystemInformationBlockType21-r14 ::= SEQUENCE
// Extensible
type Systeminformationblocktype21R14 struct {
	SlV2xConfigcommonR14     *SlV2xConfigcommonR14
	Latenoncriticalextension *utils.OCTETSTRING
}
