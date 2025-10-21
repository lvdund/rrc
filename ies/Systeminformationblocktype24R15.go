package ies

import "rrc/utils"

// SystemInformationBlockType24-r15 ::= SEQUENCE
// Extensible
type Systeminformationblocktype24R15 struct {
	CarrierfreqlistnrR15     *CarrierfreqlistnrR15
	TReselectionnrR15        TReselection
	TReselectionnrSfR15      *Speedstatescalefactors
	Latenoncriticalextension *utils.OCTETSTRING
}
