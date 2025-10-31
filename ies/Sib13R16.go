package ies

import "rrc/utils"

// SIB13-r16 ::= SEQUENCE
// Extensible
type Sib13R16 struct {
	SlV2xConfigcommonR16     utils.OCTETSTRING
	Dummy                    utils.OCTETSTRING
	TddConfigR16             utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
}
