package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest-v920-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestV920Ies struct {
	Concurrprepcdma2000HrpdR9 *bool
	Noncriticalextension      *HandoverfromeutrapreparationrequestV1020Ies
}
