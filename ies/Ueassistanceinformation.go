package ies

import "rrc/utils"

// UEAssistanceInformation-IEs ::= SEQUENCE
type Ueassistanceinformation struct {
	Delaybudgetreport        *Delaybudgetreport
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeassistanceinformationV1540
}
