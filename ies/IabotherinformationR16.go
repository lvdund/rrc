package ies

import "rrc/utils"

// IABOtherInformation-r16-IEs ::= SEQUENCE
// Extensible
type IabotherinformationR16 struct {
	IpInfotypeR16            *IabotherinformationR16IesIpInfotypeR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *IabotherinformationR16IesNoncriticalextension
}
