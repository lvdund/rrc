package ies

import "rrc/utils"

// SIB21-r17 ::= SEQUENCE
// Extensible
type Sib21R17 struct {
	MbsFsaiIntrafreqR17      *MbsFsaiListR17
	MbsFsaiInterfreqlistR17  *MbsFsaiInterfreqlistR17
	Latenoncriticalextension *utils.OCTETSTRING
}
