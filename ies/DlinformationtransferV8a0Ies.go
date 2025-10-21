package ies

import "rrc/utils"

// DLInformationTransfer-v8a0-IEs ::= SEQUENCE
type DlinformationtransferV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *DlinformationtransferV1610Ies
}
