package ies

import "rrc/utils"

// RLC-ParametersSidelink-r16-um-WithLongSN-Sidelink-r16 ::= ENUMERATED
type RlcParameterssidelinkR16UmWithlongsnSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	RlcParameterssidelinkR16UmWithlongsnSidelinkR16EnumeratedNothing = iota
	RlcParameterssidelinkR16UmWithlongsnSidelinkR16EnumeratedSupported
)
