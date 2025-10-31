package ies

import "rrc/utils"

// SL-Parameters-v1540-sl-LowT2min-r15 ::= ENUMERATED
type SlParametersV1540SlLowt2minR15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1540SlLowt2minR15EnumeratedNothing = iota
	SlParametersV1540SlLowt2minR15EnumeratedSupported
)
