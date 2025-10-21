package ies

import "rrc/utils"

// UEInformationRequest-v930-IEs ::= SEQUENCE
type UeinformationrequestV930Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeinformationrequestV1020Ies
}
