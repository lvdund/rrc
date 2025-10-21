package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest-v1020-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestV1020Ies struct {
	DualrxtxredirectindicatorR10    *utils.ENUMERATED
	Redirectcarriercdma20001xrttR10 *Carrierfreqcdma2000
	Noncriticalextension            *interface{}
}
