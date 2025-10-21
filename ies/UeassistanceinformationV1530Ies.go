package ies

import "rrc/utils"

// UEAssistanceInformation-v1530-IEs ::= SEQUENCE
type UeassistanceinformationV1530Ies struct {
	SpsAssistanceinformationV1530 *interface{}
	Noncriticalextension          *UeassistanceinformationV1610Ies
}
