package ies

import "rrc/utils"

// SidelinkParametersEUTRA-r16 ::= SEQUENCE
// Extensible
type SidelinkparameterseutraR16 struct {
	SlParameterseutra1R16             *utils.OCTETSTRING
	SlParameterseutra2R16             *utils.OCTETSTRING
	SlParameterseutra3R16             *utils.OCTETSTRING
	SupportedbandlistsidelinkeutraR16 *[]BandsidelinkeutraR16 `lb:1,ub:maxBandsEUTRA`
}
