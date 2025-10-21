package ies

import "rrc/utils"

// SystemInformationBlockType9 ::= SEQUENCE
// Extensible
type Systeminformationblocktype9 struct {
	HnbName                  *utils.OCTETSTRING
	Latenoncriticalextension *utils.OCTETSTRING
}
