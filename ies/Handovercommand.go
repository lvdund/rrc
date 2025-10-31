package ies

import "rrc/utils"

// HandoverCommand-IEs ::= SEQUENCE
type Handovercommand struct {
	Handovercommandmessage utils.OCTETSTRING
	Noncriticalextension   *HandovercommandIesNoncriticalextension
}
