package ies

import "rrc/utils"

// LAA-Parameters-v1430-uss-BlindDecodingReduction-r14 ::= ENUMERATED
type LaaParametersV1430UssBlinddecodingreductionR14 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1430UssBlinddecodingreductionR14EnumeratedNothing = iota
	LaaParametersV1430UssBlinddecodingreductionR14EnumeratedSupported
)
