package ies

import "rrc/utils"

// PhyLayerParameters-v1450 ::= SEQUENCE
type PhylayerparametersV1450 struct {
	CeSrsEnhancementwithoutcomb4R14 *utils.ENUMERATED
	CrsLessdwptsR14                 *utils.ENUMERATED
}
