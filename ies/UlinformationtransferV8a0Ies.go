package ies

import "rrc/utils"

// ULInformationTransfer-v8a0-IEs ::= SEQUENCE
type UlinformationtransferV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
