package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest-v890-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestV890Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *HandoverfromeutrapreparationrequestV920Ies
}
