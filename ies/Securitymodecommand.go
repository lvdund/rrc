package ies

import "rrc/utils"

// SecurityModeCommand-IEs ::= SEQUENCE
type Securitymodecommand struct {
	Securityconfigsmc        Securityconfigsmc
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SecuritymodecommandIesNoncriticalextension
}
