package ies

import "rrc/utils"

// UEInformationResponse-v1700-IEs ::= SEQUENCE
type UeinformationresponseV1700 struct {
	SuccesshoReportR17       *SuccesshoReportR17
	ConnestfailreportlistR17 *ConnestfailreportlistR17
	CoarselocationinfoR17    *utils.OCTETSTRING
	Noncriticalextension     *UeinformationresponseV1700IesNoncriticalextension
}
