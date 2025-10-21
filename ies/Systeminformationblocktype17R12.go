package ies

import "rrc/utils"

// SystemInformationBlockType17-r12 ::= SEQUENCE
// Extensible
type Systeminformationblocktype17R12 struct {
	WlanOffloadinfoperplmnListR12 *interface{}
	Latenoncriticalextension      *utils.OCTETSTRING
}
