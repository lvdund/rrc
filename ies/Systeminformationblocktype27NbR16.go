package ies

import "rrc/utils"

// SystemInformationBlockType27-NB-r16 ::= SEQUENCE
// Extensible
type Systeminformationblocktype27NbR16 struct {
	CarrierfreqlisteutraR16  *CarrierfreqlisteutraNbR16
	CarrierfreqslistgeranR16 *CarrierfreqslistgeranNbR16
	Latenoncriticalextension *utils.OCTETSTRING
}
