package ies

import "rrc/utils"

// UEInformationRequest-v1250-IEs ::= SEQUENCE
type UeinformationrequestV1250Ies struct {
	MobilityhistoryreportreqR12 *utils.ENUMERATED
	Noncriticalextension        *UeinformationrequestV1530Ies
}
