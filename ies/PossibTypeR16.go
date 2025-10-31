package ies

import "rrc/utils"

// PosSIB-Type-r16 ::= SEQUENCE
// Extensible
type PossibTypeR16 struct {
	EncryptedR16  *utils.BOOLEAN
	GnssIdR16     *GnssIdR16
	SbasIdR16     *SbasIdR16
	PossibtypeR16 PossibTypeR16PossibtypeR16
	AreascopeR16  *utils.BOOLEAN
}
