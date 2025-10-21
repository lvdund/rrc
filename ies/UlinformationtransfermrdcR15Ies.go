package ies

import "rrc/utils"

// ULInformationTransferMRDC-r15-IEs ::= SEQUENCE
type UlinformationtransfermrdcR15Ies struct {
	UlDcchMessagenrR15       *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
