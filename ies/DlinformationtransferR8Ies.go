package ies

import "rrc/utils"

// DLInformationTransfer-r8-IEs ::= SEQUENCE
type DlinformationtransferR8Ies struct {
	Dedicatedinfotype    interface{}
	Noncriticalextension *DlinformationtransferV8a0Ies
}
