package ies

import "rrc/utils"

// SecurityModeCommand-v8a0-IEs ::= SEQUENCE
type SecuritymodecommandV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SecuritymodecommandV8a0IesNoncriticalextension
}
