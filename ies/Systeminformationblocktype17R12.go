package ies

import "rrc/utils"

// SystemInformationBlockType17-r12 ::= SEQUENCE
// Extensible
type Systeminformationblocktype17R12 struct {
	WlanOffloadinfoperplmnListR12 *[]WlanOffloadinfoperplmnR12 `lb:1,ub:maxPLMNR11`
	Latenoncriticalextension      *utils.OCTETSTRING
}
