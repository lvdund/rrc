package ies

import "rrc/utils"

// ULInformationTransferMRDC-IEs ::= SEQUENCE
type Ulinformationtransfermrdc struct {
	UlDcchMessagenr          *utils.OCTETSTRING
	UlDcchMessageeutra       *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlinformationtransfermrdcIesNoncriticalextension
}
