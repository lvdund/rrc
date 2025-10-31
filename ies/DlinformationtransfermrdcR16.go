package ies

import "rrc/utils"

// DLInformationTransferMRDC-r16-IEs ::= SEQUENCE
type DlinformationtransfermrdcR16 struct {
	DlDcchMessagenrR16       *utils.OCTETSTRING
	DlDcchMessageeutraR16    *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *DlinformationtransfermrdcR16IesNoncriticalextension
}
