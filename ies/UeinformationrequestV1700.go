package ies

import "rrc/utils"

// UEInformationRequest-v1700-IEs ::= SEQUENCE
type UeinformationrequestV1700 struct {
	SuccesshoReportreqR17    *utils.BOOLEAN
	CoarselocationrequestR17 *utils.BOOLEAN
	Noncriticalextension     *UeinformationrequestV1700IesNoncriticalextension
}
