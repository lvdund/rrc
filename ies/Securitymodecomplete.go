package ies

import "rrc/utils"

// SecurityModeComplete-IEs ::= SEQUENCE
type Securitymodecomplete struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SecuritymodecompleteIesNoncriticalextension
}
