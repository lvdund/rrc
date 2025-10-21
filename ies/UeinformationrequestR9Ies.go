package ies

import "rrc/utils"

// UEInformationRequest-r9-IEs ::= SEQUENCE
type UeinformationrequestR9Ies struct {
	RachReportreqR9      bool
	RlfReportreqR9       bool
	Noncriticalextension *UeinformationrequestV930Ies
}
