package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-1XRTT-v1020-e-CSFB-dual-1XRTT-r10 ::= ENUMERATED
type IratParameterscdma20001xrttV1020ECsfbDual1xrttR10 struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma20001xrttV1020ECsfbDual1xrttR10EnumeratedNothing = iota
	IratParameterscdma20001xrttV1020ECsfbDual1xrttR10EnumeratedSupported
)
