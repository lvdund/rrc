package ies

import "rrc/utils"

// SecurityModeCommand-r8-IEs ::= SEQUENCE
type SecuritymodecommandR8Ies struct {
	Securityconfigsmc    Securityconfigsmc
	Noncriticalextension *SecuritymodecommandV8a0Ies
}
