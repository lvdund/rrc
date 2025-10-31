package ies

import "rrc/utils"

// SN-FieldLengthUM ::= ENUMERATED
type SnFieldlengthum struct {
	Value utils.ENUMERATED
}

const (
	SnFieldlengthumEnumeratedNothing = iota
	SnFieldlengthumEnumeratedSize6
	SnFieldlengthumEnumeratedSize12
)
