package ies

import "rrc/utils"

// RLC-ParametersSidelink-r16-am-WithLongSN-Sidelink-r16 ::= ENUMERATED
type RlcParameterssidelinkR16AmWithlongsnSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	RlcParameterssidelinkR16AmWithlongsnSidelinkR16EnumeratedNothing = iota
	RlcParameterssidelinkR16AmWithlongsnSidelinkR16EnumeratedSupported
)
