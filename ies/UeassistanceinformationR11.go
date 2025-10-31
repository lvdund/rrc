package ies

import "rrc/utils"

// UEAssistanceInformation-r11-IEs ::= SEQUENCE
type UeassistanceinformationR11 struct {
	PowerprefindicationR11   *UeassistanceinformationR11IesPowerprefindicationR11
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeassistanceinformationV1430
}
