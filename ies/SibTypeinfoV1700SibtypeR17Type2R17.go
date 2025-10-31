package ies

import "rrc/utils"

// SIB-TypeInfo-v1700-sibType-r17-type2-r17 ::= SEQUENCE
// Extensible
type SibTypeinfoV1700SibtypeR17Type2R17 struct {
	PossibtypeR17 SibTypeinfoV1700SibtypeR17Type2R17PossibtypeR17
	EncryptedR17  *utils.BOOLEAN
	GnssIdR17     *GnssIdR16
	SbasIdR17     *SbasIdR16
}
