package ies

import "rrc/utils"

// SL-Parameters-v1540-sl-64QAM-Rx-r15 ::= ENUMERATED
type SlParametersV1540Sl64qamRxR15 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1540Sl64qamRxR15EnumeratedNothing = iota
	SlParametersV1540Sl64qamRxR15EnumeratedSupported
)
