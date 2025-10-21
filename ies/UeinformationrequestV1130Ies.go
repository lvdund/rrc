package ies

import "rrc/utils"

// UEInformationRequest-v1130-IEs ::= SEQUENCE
type UeinformationrequestV1130Ies struct {
	ConnestfailreportreqR11 *utils.ENUMERATED
	Noncriticalextension    *UeinformationrequestV1250Ies
}
