package ies

import "rrc/utils"

// UEInformationRequest-r16-IEs ::= SEQUENCE
type UeinformationrequestR16 struct {
	IdlemodemeasurementreqR16   *utils.BOOLEAN
	LogmeasreportreqR16         *utils.BOOLEAN
	ConnestfailreportreqR16     *utils.BOOLEAN
	RaReportreqR16              *utils.BOOLEAN
	RlfReportreqR16             *utils.BOOLEAN
	MobilityhistoryreportreqR16 *utils.BOOLEAN
	Latenoncriticalextension    *utils.OCTETSTRING
	Noncriticalextension        *UeinformationrequestV1700
}
