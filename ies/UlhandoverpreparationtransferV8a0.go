package ies

import "rrc/utils"

// ULHandoverPreparationTransfer-v8a0-IEs ::= SEQUENCE
type UlhandoverpreparationtransferV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UlhandoverpreparationtransferV8a0IesNoncriticalextension
}
