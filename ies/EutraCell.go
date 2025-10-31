package ies

// EUTRA-Cell ::= SEQUENCE
type EutraCell struct {
	Cellindexeutra       EutraCellindex
	Physcellid           EutraPhyscellid
	Cellindividualoffset EutraQOffsetrange
}
