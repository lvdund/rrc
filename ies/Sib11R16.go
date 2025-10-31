package ies

import "rrc/utils"

// SIB11-r16 ::= SEQUENCE
// Extensible
type Sib11R16 struct {
	MeasidleconfigsibR16     *MeasidleconfigsibR16
	Latenoncriticalextension *utils.OCTETSTRING
}
