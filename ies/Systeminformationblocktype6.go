package ies

import "rrc/utils"

// SystemInformationBlockType6 ::= SEQUENCE
// Extensible
type Systeminformationblocktype6 struct {
	CarrierfreqlistutraFdd   *CarrierfreqlistutraFdd
	CarrierfreqlistutraTdd   *CarrierfreqlistutraTdd
	TReselectionutra         TReselection
	TReselectionutraSf       *Speedstatescalefactors
	Latenoncriticalextension *utils.OCTETSTRING
}
