package ies

import "rrc/utils"

// UEAssistanceInformation-r11-IEs ::= SEQUENCE
type UeassistanceinformationR11Ies struct {
	PowerprefindicationR11   *utils.ENUMERATED
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeassistanceinformationV1430Ies
}
