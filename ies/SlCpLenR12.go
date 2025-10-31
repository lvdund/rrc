package ies

import "rrc/utils"

// SL-CP-Len-r12 ::= ENUMERATED
type SlCpLenR12 struct {
	Value utils.ENUMERATED
}

const (
	SlCpLenR12EnumeratedNothing = iota
	SlCpLenR12EnumeratedNormal
	SlCpLenR12EnumeratedExtended
)
