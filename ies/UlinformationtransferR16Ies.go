package ies

import "rrc/utils"

// ULInformationTransfer-r16-IEs ::= SEQUENCE
type UlinformationtransferR16Ies struct {
	DedicatedinfotypeR16 *interface{}
	Dedicatedinfof1cR16  *Dedicatedinfof1cR16
	Noncriticalextension *UlinformationtransferV8a0Ies
}
