package ies

import "rrc/utils"

// UEAssistanceInformation-v1450-IEs ::= SEQUENCE
type UeassistanceinformationV1450Ies struct {
	OverheatingassistanceR14 *OverheatingassistanceR14
	Noncriticalextension     *UeassistanceinformationV1530Ies
}
