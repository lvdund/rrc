package ies

import "rrc/utils"

// RF-Parameters-v1430-diffFallbackCombReport-r14 ::= ENUMERATED
type RfParametersV1430DifffallbackcombreportR14 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1430DifffallbackcombreportR14EnumeratedNothing = iota
	RfParametersV1430DifffallbackcombreportR14EnumeratedSupported
)
