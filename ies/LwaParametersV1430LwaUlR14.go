package ies

import "rrc/utils"

// LWA-Parameters-v1430-lwa-UL-r14 ::= ENUMERATED
type LwaParametersV1430LwaUlR14 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersV1430LwaUlR14EnumeratedNothing = iota
	LwaParametersV1430LwaUlR14EnumeratedSupported
)
