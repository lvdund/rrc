package ies

import "rrc/utils"

// ULInformationTransferIRAT-r16-IEs ::= SEQUENCE
type UlinformationtransferiratR16Ies struct {
	UlDcchMessagenrR16       *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
