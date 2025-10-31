package ies

import "rrc/utils"

// RLC-Parameters-um-WithShortSN ::= ENUMERATED
type RlcParametersUmWithshortsn struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersUmWithshortsnEnumeratedNothing = iota
	RlcParametersUmWithshortsnEnumeratedSupported
)
