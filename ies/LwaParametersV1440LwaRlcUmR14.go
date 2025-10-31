package ies

import "rrc/utils"

// LWA-Parameters-v1440-lwa-RLC-UM-r14 ::= ENUMERATED
type LwaParametersV1440LwaRlcUmR14 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersV1440LwaRlcUmR14EnumeratedNothing = iota
	LwaParametersV1440LwaRlcUmR14EnumeratedSupported
)
