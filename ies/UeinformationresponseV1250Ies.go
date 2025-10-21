package ies

import "rrc/utils"

// UEInformationResponse-v1250-IEs ::= SEQUENCE
type UeinformationresponseV1250Ies struct {
	MobilityhistoryreportR12 *MobilityhistoryreportR12
	Noncriticalextension     *UeinformationresponseV1530Ies
}
