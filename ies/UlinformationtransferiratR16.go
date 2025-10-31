package ies

import "rrc/utils"

// ULInformationTransferIRAT-r16-IEs ::= SEQUENCE
type UlinformationtransferiratR16 struct {
	UlDcchMessageeutraR16    *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlinformationtransferiratR16IesNoncriticalextension
}
