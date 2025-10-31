package ies

import "rrc/utils"

// RLC-Parameters-am-WithShortSN ::= ENUMERATED
type RlcParametersAmWithshortsn struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersAmWithshortsnEnumeratedNothing = iota
	RlcParametersAmWithshortsnEnumeratedSupported
)
