package ies

import "rrc/utils"

// SystemInformationBlockType9 ::= SEQUENCE
// Extensible
type Systeminformationblocktype9 struct {
	HnbName                  *utils.OCTETSTRING `lb:1,ub:48`
	Latenoncriticalextension *utils.OCTETSTRING
}
