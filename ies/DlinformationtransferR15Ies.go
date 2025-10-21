package ies

import "rrc/utils"

// DLInformationTransfer-r15-IEs ::= SEQUENCE
type DlinformationtransferR15Ies struct {
	DedicatedinfotypeR15 *interface{}
	TimereferenceinfoR15 *TimereferenceinfoR15
	Noncriticalextension *DlinformationtransferV8a0Ies
}
