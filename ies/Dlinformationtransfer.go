package ies

import "rrc/utils"

// DLInformationTransfer-IEs ::= SEQUENCE
type Dlinformationtransfer struct {
	DedicatednasMessage      *DedicatednasMessage
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *DlinformationtransferV1610
}
