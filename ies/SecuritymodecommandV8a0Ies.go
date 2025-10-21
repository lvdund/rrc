package ies

import "rrc/utils"

// SecurityModeCommand-v8a0-IEs ::= SEQUENCE
type SecuritymodecommandV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
