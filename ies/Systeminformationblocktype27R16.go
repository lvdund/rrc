package ies

import "rrc/utils"

// SystemInformationBlockType27-r16 ::= SEQUENCE
// Extensible
type Systeminformationblocktype27R16 struct {
	CarrierfreqlistnbiotR16  *CarrierfreqlistnbiotR16
	Latenoncriticalextension *utils.OCTETSTRING
}
