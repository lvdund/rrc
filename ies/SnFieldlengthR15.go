package ies

import "rrc/utils"

// SN-FieldLength-r15 ::= ENUMERATED
type SnFieldlengthR15 struct {
	Value utils.ENUMERATED
}

const (
	SnFieldlengthR15Size5     = 0
	SnFieldlengthR15Size10    = 1
	SnFieldlengthR15Size16R15 = 2
)
