package ies

import "rrc/utils"

// HandoverCommand-r8-IEs ::= SEQUENCE
type HandovercommandR8 struct {
	Handovercommandmessage utils.OCTETSTRING
	Noncriticalextension   *HandovercommandR8IesNoncriticalextension
}
