package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest-v890-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestV890 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *HandoverfromeutrapreparationrequestV920
}
