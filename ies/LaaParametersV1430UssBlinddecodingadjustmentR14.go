package ies

import "rrc/utils"

// LAA-Parameters-v1430-uss-BlindDecodingAdjustment-r14 ::= ENUMERATED
type LaaParametersV1430UssBlinddecodingadjustmentR14 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1430UssBlinddecodingadjustmentR14EnumeratedNothing = iota
	LaaParametersV1430UssBlinddecodingadjustmentR14EnumeratedSupported
)
