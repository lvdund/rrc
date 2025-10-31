package ies

import "rrc/utils"

// SN-FieldLength ::= ENUMERATED
type SnFieldlength struct {
	Value utils.ENUMERATED
}

const (
	SnFieldlengthEnumeratedNothing = iota
	SnFieldlengthEnumeratedSize5
	SnFieldlengthEnumeratedSize10
)
