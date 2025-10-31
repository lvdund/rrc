package ies

import "rrc/utils"

// ULInformationTransfer-NB-r13-IEs ::= SEQUENCE
type UlinformationtransferNbR13 struct {
	DedicatedinfonasR13      Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlinformationtransferNbR13IesNoncriticalextension
}
