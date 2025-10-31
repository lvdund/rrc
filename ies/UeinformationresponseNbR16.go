package ies

import "rrc/utils"

// UEInformationResponse-NB-r16-IEs ::= SEQUENCE
type UeinformationresponseNbR16 struct {
	RachReportR16            *RachReportNbR16
	RlfReportR16             *RlfReportNbR16
	AnrMeasreportR16         *AnrMeasreportNbR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeinformationresponseNbR16IesNoncriticalextension
}
