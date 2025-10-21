package ies

import "rrc/utils"

// HandoverCommand-r8-IEs ::= SEQUENCE
type HandovercommandR8Ies struct {
	Handovercommandmessage utils.OCTETSTRING
	Noncriticalextension   *interface{}
}
