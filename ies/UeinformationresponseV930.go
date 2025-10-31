package ies

import "rrc/utils"

// UEInformationResponse-v930-IEs ::= SEQUENCE
type UeinformationresponseV930 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeinformationresponseV1020
}
