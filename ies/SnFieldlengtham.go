package ies

import "rrc/utils"

// SN-FieldLengthAM ::= ENUMERATED
type SnFieldlengtham struct {
	Value utils.ENUMERATED
}

const (
	SnFieldlengthamEnumeratedNothing = iota
	SnFieldlengthamEnumeratedSize12
	SnFieldlengthamEnumeratedSize18
)
