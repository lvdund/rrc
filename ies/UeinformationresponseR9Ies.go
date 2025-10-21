package ies

import "rrc/utils"

// UEInformationResponse-r9-IEs ::= SEQUENCE
type UeinformationresponseR9Ies struct {
	RachReportR9         *RachReportR16
	RlfReportR9          *RlfReportR9
	Noncriticalextension *UeinformationresponseV930Ies
}
