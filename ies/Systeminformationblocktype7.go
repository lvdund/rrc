package ies

import "rrc/utils"

// SystemInformationBlockType7 ::= SEQUENCE
// Extensible
type Systeminformationblocktype7 struct {
	TReselectiongeran        TReselection
	TReselectiongeranSf      *Speedstatescalefactors
	Carrierfreqsinfolist     *Carrierfreqsinfolistgeran
	Latenoncriticalextension *utils.OCTETSTRING
}
