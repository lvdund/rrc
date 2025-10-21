package ies

import "rrc/utils"

// UEInformationRequest-v1020-IEs ::= SEQUENCE
type UeinformationrequestV1020Ies struct {
	LogmeasreportreqR10  *utils.ENUMERATED
	Noncriticalextension *UeinformationrequestV1130Ies
}
