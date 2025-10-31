package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest-v920-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestV920 struct {
	Concurrprepcdma2000HrpdR9 *utils.BOOLEAN
	Noncriticalextension      *HandoverfromeutrapreparationrequestV1020
}
