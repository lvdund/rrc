package ies

// PosSIB-Type-r15 ::= SEQUENCE
// Extensible
type PossibTypeR15 struct {
	EncryptedR15  *bool
	GnssIdR15     *GnssIdR15
	SbasIdR15     *SbasIdR15
	PossibtypeR15 PossibTypeR15PossibtypeR15
}
