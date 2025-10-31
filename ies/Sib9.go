package ies

import "rrc/utils"

// SIB9 ::= SEQUENCE
// Extensible
type Sib9 struct {
	Timeinfo                 *Sib9Timeinfo
	Latenoncriticalextension *utils.OCTETSTRING
	ReferencetimeinfoR16     *ReferencetimeinfoR16 `ext`
}
