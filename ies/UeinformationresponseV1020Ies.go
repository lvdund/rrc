package ies

import "rrc/utils"

// UEInformationResponse-v1020-IEs ::= SEQUENCE
type UeinformationresponseV1020Ies struct {
	LogmeasreportR10     *LogmeasreportR10
	Noncriticalextension *UeinformationresponseV1130Ies
}
