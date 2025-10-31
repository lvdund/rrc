package ies

import "rrc/utils"

// DLInformationTransfer-v1700-IEs ::= SEQUENCE
type DlinformationtransferV1700 struct {
	Dedicatedinfof1cR17  *Dedicatedinfof1cR17
	RxtxtimediffGnbR17   *RxtxtimediffR17
	TaPdcR17             *DlinformationtransferV1700IesTaPdcR17
	Sib9fallbackR17      *utils.BOOLEAN
	Noncriticalextension *DlinformationtransferV1700IesNoncriticalextension
}
