package ies

import "rrc/utils"

// ULInformationTransfer-IEs ::= SEQUENCE
type Ulinformationtransfer struct {
	DedicatednasMessage      *DedicatednasMessage
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlinformationtransferV1700
}
