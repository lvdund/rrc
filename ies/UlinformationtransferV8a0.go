package ies

import "rrc/utils"

// ULInformationTransfer-v8a0-IEs ::= SEQUENCE
type UlinformationtransferV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlinformationtransferV8a0IesNoncriticalextension
}
