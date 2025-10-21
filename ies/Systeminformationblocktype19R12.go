package ies

import "rrc/utils"

// SystemInformationBlockType19-r12 ::= SEQUENCE
// Extensible
type Systeminformationblocktype19R12 struct {
	DiscconfigR12            *interface{}
	DiscinterfreqlistR12     *SlCarrierfreqinfolistR12
	Latenoncriticalextension *utils.OCTETSTRING
}
