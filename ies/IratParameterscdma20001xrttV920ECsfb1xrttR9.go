package ies

import "rrc/utils"

// IRAT-ParametersCDMA2000-1XRTT-v920-e-CSFB-1XRTT-r9 ::= ENUMERATED
type IratParameterscdma20001xrttV920ECsfb1xrttR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParameterscdma20001xrttV920ECsfb1xrttR9EnumeratedNothing = iota
	IratParameterscdma20001xrttV920ECsfb1xrttR9EnumeratedSupported
)
