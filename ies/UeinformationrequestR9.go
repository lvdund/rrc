package ies

import "rrc/utils"

// UEInformationRequest-r9-IEs ::= SEQUENCE
type UeinformationrequestR9 struct {
	RachReportreqR9      utils.BOOLEAN
	RlfReportreqR9       utils.BOOLEAN
	Noncriticalextension *UeinformationrequestV930
}
