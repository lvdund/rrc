package ies

import "rrc/utils"

// LAA-Parameters-v1430-crossCarrierSchedulingLAA-UL-r14 ::= ENUMERATED
type LaaParametersV1430CrosscarrierschedulinglaaUlR14 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1430CrosscarrierschedulinglaaUlR14EnumeratedNothing = iota
	LaaParametersV1430CrosscarrierschedulinglaaUlR14EnumeratedSupported
)
