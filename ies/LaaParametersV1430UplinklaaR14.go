package ies

import "rrc/utils"

// LAA-Parameters-v1430-uplinkLAA-r14 ::= ENUMERATED
type LaaParametersV1430UplinklaaR14 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1430UplinklaaR14EnumeratedNothing = iota
	LaaParametersV1430UplinklaaR14EnumeratedSupported
)
