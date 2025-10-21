package ies

import "rrc/utils"

// UEInformationResponse-v1130-IEs ::= SEQUENCE
type UeinformationresponseV1130Ies struct {
	ConnestfailreportR11 *ConnestfailreportR11
	Noncriticalextension *UeinformationresponseV1250Ies
}
