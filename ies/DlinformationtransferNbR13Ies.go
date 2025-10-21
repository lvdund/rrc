package ies

import "rrc/utils"

// DLInformationTransfer-NB-r13-IEs ::= SEQUENCE
type DlinformationtransferNbR13Ies struct {
	DedicatedinfonasR13      Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
