package ies

import "rrc/utils"

// SIB10-r16 ::= SEQUENCE
// Extensible
type Sib10R16 struct {
	HrnnListR16              *HrnnListR16
	Latenoncriticalextension *utils.OCTETSTRING
}
