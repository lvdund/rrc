package ies

import "rrc/utils"

// UEInformationRequest-NB-r16-IEs ::= SEQUENCE
type UeinformationrequestNbR16 struct {
	RachReportreqR16         utils.BOOLEAN
	RlfReportreqR16          utils.BOOLEAN
	AnrReportreqR16          utils.BOOLEAN
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeinformationrequestNbR16IesNoncriticalextension
}
