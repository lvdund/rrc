package ies

import "rrc/utils"

// SystemInformationBlockType5-NB-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype5NbR13 struct {
	InterfreqcarrierfreqlistR13 InterfreqcarrierfreqlistNbR13
	TReselectionR13             TReselectionNbR13
	Latenoncriticalextension    *utils.OCTETSTRING
}
