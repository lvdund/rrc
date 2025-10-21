package ies

import "rrc/utils"

// SecurityModeComplete-v8a0-IEs ::= SEQUENCE
type SecuritymodecompleteV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
