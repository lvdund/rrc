package ies

import "rrc/utils"

// SN-FieldLength ::= ENUMERATED
type SnFieldlength struct {
	Value utils.ENUMERATED
}

const (
	SnFieldlengthSize5  = 0
	SnFieldlengthSize10 = 1
)
