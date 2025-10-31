package ies

import "rrc/utils"

// UEInformationRequest-v930-IEs ::= SEQUENCE
type UeinformationrequestV930 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeinformationrequestV1020
}
