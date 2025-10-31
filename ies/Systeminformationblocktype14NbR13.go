package ies

import "rrc/utils"

// SystemInformationBlockType14-NB-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype14NbR13 struct {
	AbParamR13               *Systeminformationblocktype14NbR13AbParamR13
	Latenoncriticalextension *utils.OCTETSTRING
}
