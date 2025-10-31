package ies

import "rrc/utils"

// RLC-Parameters-um-WithLongSN ::= ENUMERATED
type RlcParametersUmWithlongsn struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersUmWithlongsnEnumeratedNothing = iota
	RlcParametersUmWithlongsnEnumeratedSupported
)
