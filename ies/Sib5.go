package ies

import "rrc/utils"

// SIB5 ::= SEQUENCE
// Extensible
type Sib5 struct {
	Carrierfreqlisteutra         *Carrierfreqlisteutra
	TReselectioneutra            TReselection
	TReselectioneutraSf          *Speedstatescalefactors
	Latenoncriticalextension     *utils.OCTETSTRING
	CarrierfreqlisteutraV1610    *CarrierfreqlisteutraV1610 `ext`
	CarrierfreqlisteutraV1700    *CarrierfreqlisteutraV1700 `ext`
	IdlemodemeasvoicefallbackR17 *utils.BOOLEAN             `ext`
}
