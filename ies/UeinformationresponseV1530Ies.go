package ies

import "rrc/utils"

// UEInformationResponse-v1530-IEs ::= SEQUENCE
type UeinformationresponseV1530Ies struct {
	MeasresultlistidleR15   *MeasresultlistidleR15
	FlightpathinforeportR15 *FlightpathinforeportR15
	Noncriticalextension    *UeinformationresponseV1610Ies
}
