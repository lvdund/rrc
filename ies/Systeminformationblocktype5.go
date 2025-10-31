package ies

import "rrc/utils"

// SystemInformationBlockType5 ::= SEQUENCE
// Extensible
type Systeminformationblocktype5 struct {
	Interfreqcarrierfreqlist Interfreqcarrierfreqlist
	Latenoncriticalextension *utils.OCTETSTRING
}
