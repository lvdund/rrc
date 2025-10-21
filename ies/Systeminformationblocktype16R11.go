package ies

import "rrc/utils"

// SystemInformationBlockType16-r11 ::= SEQUENCE
// Extensible
type Systeminformationblocktype16R11 struct {
	TimeinfoR11              *interface{}
	Latenoncriticalextension *utils.OCTETSTRING
}
