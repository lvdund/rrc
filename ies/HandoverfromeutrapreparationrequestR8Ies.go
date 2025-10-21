package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest-r8-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestR8Ies struct {
	Cdma2000Type         Cdma2000Type
	Rand                 *RandCdma2000
	Mobilityparameters   *Mobilityparameterscdma2000
	Noncriticalextension *HandoverfromeutrapreparationrequestV890Ies
}
