package ies

import "rrc/utils"

// UEInformationRequest-NB-r16-IEs ::= SEQUENCE
type UeinformationrequestNbR16Ies struct {
	RachReportreqR16         bool
	RlfReportreqR16          bool
	AnrReportreqR16          bool
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
