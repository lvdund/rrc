package ies

import "rrc/utils"

// SecurityModeComplete-v8a0-IEs ::= SEQUENCE
type SecuritymodecompleteV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SecuritymodecompleteV8a0IesNoncriticalextension
}
