package ies

import "rrc/utils"

// SL-CP-Len-r12 ::= ENUMERATED
type SlCpLenR12 struct {
	Value utils.ENUMERATED
}

const (
	SlCpLenR12Normal   = 0
	SlCpLenR12Extended = 1
)
