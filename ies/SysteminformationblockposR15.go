package ies

import "rrc/utils"

// SystemInformationBlockPos-r15 ::= SEQUENCE
// Extensible
type SysteminformationblockposR15 struct {
	AssistancedatasibElementR15 utils.OCTETSTRING
	Latenoncriticalextension    *utils.OCTETSTRING
}
