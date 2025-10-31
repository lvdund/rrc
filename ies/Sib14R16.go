package ies

import "rrc/utils"

// SIB14-r16 ::= SEQUENCE
// Extensible
type Sib14R16 struct {
	SlV2xConfigcommonextR16  utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
}
