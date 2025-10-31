package ies

import "rrc/utils"

// RLC-Parameters-r12-extended-RLC-LI-Field-r12 ::= ENUMERATED
type RlcParametersR12ExtendedRlcLiFieldR12 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersR12ExtendedRlcLiFieldR12EnumeratedNothing = iota
	RlcParametersR12ExtendedRlcLiFieldR12EnumeratedSupported
)
