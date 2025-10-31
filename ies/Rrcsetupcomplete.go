package ies

import "rrc/utils"

// RRCSetupComplete-IEs ::= SEQUENCE
type Rrcsetupcomplete struct {
	SelectedplmnIdentity     utils.INTEGER `lb:0,ub:maxPLMN`
	Registeredamf            *Registeredamf
	GuamiType                *RrcsetupcompleteIesGuamiType
	SNssaiList               *[]SNssai `lb:1,ub:maxNrofSNssai`
	DedicatednasMessage      DedicatednasMessage
	Ng5gSTmsiValue           *RrcsetupcompleteIesNg5gSTmsiValue
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcsetupcompleteV1610
}
