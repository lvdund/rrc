package ies

import "rrc/utils"

// ULInformationTransferMRDC-r15-IEs ::= SEQUENCE
type UlinformationtransfermrdcR15 struct {
	UlDcchMessagenrR15       *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlinformationtransfermrdcR15IesNoncriticalextension
}
