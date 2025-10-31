package ies

import "rrc/utils"

// SIBpos-r16 ::= SEQUENCE
// Extensible
type SibposR16 struct {
	AssistancedatasibElementR16 utils.OCTETSTRING
	Latenoncriticalextension    *utils.OCTETSTRING
}
